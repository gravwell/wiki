package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

const (
	baseTimeout            = 10 * time.Second
	mdExt           string = `.md`
	maxMarkdownSize        = 0x800000 //8MB which is bonkers
)

var (
	lp     = flag.Uint("listen-port", 3001, "listening port")
	la     = flag.String("listen-address", "", "listening address")
	sdir   = flag.String("host-dir", ".", "Root directory for files")
	blink  = flag.String("base-link", "#!", "Base value to return for all links")
	fdebug = flag.Bool("debug", false, "Output a ton of info while running")
)

type server struct {
	sync.Mutex
	http.Server
	blink string //the base link anchor
	dir   string
	links map[string][]ref
}

func main() {
	flag.Parse()
	if *lp == 0 || *lp > 0xffff {
		log.Fatal("invalid listening port, must be a valid TCP port")
	}
	d := filepath.Clean(*sdir)
	if fi, err := os.Stat(d); err != nil {
		log.Fatalf("failed to stat %s: %v\n", d, err)
	} else if !fi.IsDir() {
		log.Fatalf("failed to stat %s\n", d)
	}

	//get the HTTP server up and rolling
	srv := &server{
		Server: http.Server{
			Addr:         fmt.Sprintf("%v:%d", *la, *lp),
			ReadTimeout:  baseTimeout,
			WriteTimeout: baseTimeout,
		},
		dir:   *sdir,
		blink: *blink,
		links: map[string][]ref{}, //just to prevent nil-pointer errors

	}
	mx := http.NewServeMux()
	mx.Handle("/", noCacheLoggingHandler(http.FileServer(http.Dir(*sdir))))
	mx.HandleFunc("/api/search", srv.search)
	srv.Handler = mx

	if err := srv.reloadIndex(); err != nil {
		log.Fatalf("Failed to perform an initial load on the search index")
	}

	//get a signal handler up and listen for exits and reloads
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM)
	go func(s *server, sc chan os.Signal) {
		for {
			sig, ok := <-sigChan
			if !ok {
				s.Shutdown(context.Background())
				return
			}
			switch sig {
			case syscall.SIGHUP: //reload the index
				if err := s.reloadIndex(); err != nil {
					log.Printf("Failed to reload index: %v\n", err)
				}
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM:
				s.Shutdown(context.Background())
				return
			}
		}
	}(srv, sigChan)

	//actually fire up the server
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

// Wrap the given handler, adding headers to disable caching
func noCacheLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		h.ServeHTTP(w, r)
	})
}

func (s *server) reloadIndex() (err error) {

	links := make(map[string][]ref)

	//walk the directory
	err = filepath.Walk(s.dir, func(path string, info os.FileInfo, err error) error {
		//check errors and clean up paths
		if err != nil {
			return err
		}
		if path = filepath.Clean(path); filepath.Ext(path) != mdExt {
			return nil
		} else if info.Size() > maxMarkdownSize {
			log.Printf("Skipping %s due to oversized file %x > %x", path, info.Size(), maxMarkdownSize)
			return nil
		}

		// render the content
		relPath, err := filepath.Rel(s.dir, path)
		if err != nil {
			return err
		}
		return s.renderMarkdownFile(path, relPath, links)
	})
	s.Lock()
	s.links = links
	s.Unlock()
	return
}

// renderMarkdownFile reads the entire file into memory, so caller should do some sanity checking
func (s *server) renderMarkdownFile(pth, relpath string, links map[string][]ref) (err error) {
	var md []byte
	if md, err = ioutil.ReadFile(pth); err != nil {
		return
	}
	extensions := parser.CommonExtensions | parser.FencedCode | parser.AutoHeadingIDs | parser.Footnotes
	p := parser.NewWithExtensions(extensions)
	if n := p.Parse(md); n == nil {
		err = fmt.Errorf("%s is an invalid markdown file", pth)
	} else {
		cr := &cracker{
			pth:   s.blink + relpath,
			links: links,
			debug: *fdebug,
		}
		err = cr.extractLinks(n)
		if *fdebug {
			fmt.Println("Processed", relpath)
		}
	}
	return
}

func (s *server) searchTerms(v string) (r []string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(v))
	// Set the split function for the scanning operation.
	scanner.Split(scanwords)
	for scanner.Scan() {
		if txt, ok := indexableString(scanner.Text()); ok {
			if strings.ContainsAny(txt, " \t\n") {
				//process the sub strings
				var x []string
				if x, err = s.searchTerms(txt); err != nil {
					return
				} else {
					r = append(r, x...)
				}
			} else {
				r = append(r, txt)
			}
		}
	}
	err = scanner.Err()
	return
}

type search struct {
	Value string `json:"value,omitempty"`
}

type response struct {
	Links []ref `json:"links"`
}

func (r response) MarshalJSON() (bts []byte, err error) {
	if len(r.Links) == 0 {
		bts = []byte(`{"links":[]}`)
	} else {
		mr := struct {
			Links []ref `json:"links"`
		}{
			Links: r.Links,
		}
		bts, err = json.Marshal(mr)
	}
	return
}

type respError struct {
	Error string `jons:"error,omitempt"`
}

func (s *server) search(w http.ResponseWriter, r *http.Request) {
	var req search
	var resp response
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	switch r.Method {
	case http.MethodPost: //all good
	case http.MethodHead:
		return //still all good but we are done here
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		rerr := respError{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(rerr)
		return
	}
	terms, err := s.searchTerms(strings.ToLower(req.Value))
	if err != nil {
		rerr := respError{
			Error: err.Error(),
		}
		json.NewEncoder(w).Encode(rerr)
		return
	} else if len(terms) == 0 {
		rerr := respError{
			Error: `Invalid search terms`,
		}
		json.NewEncoder(w).Encode(rerr)
		return
	}
	var linkset [][]ref
	s.Lock()
	for _, term := range terms {
		if v, ok := s.links[term]; ok {
			linkset = append(linkset, v)
		}
	}
	s.Unlock()
	resp.Links = union(linkset)
	sortLinks(resp.Links)
	json.NewEncoder(w).Encode(resp)

	return
}

var empty es

type es struct{}

// this is a CRAZY expensive method for creating a union, something better?
func union(sls [][]ref) []ref {
	if len(sls) == 0 {
		return nil
	} else if len(sls) == 1 {
		return sls[0]
	}
	v := sls[0]
	sls = sls[1:]

	mp := make(map[ref]es, len(v))
	for i := range v {
		mp[v[i]] = empty
	}

	for _, sl := range sls {
		for k := range mp {
			if !inRefList(k, sl) {
				delete(mp, k)
			}
		}
	}

	r := make([]ref, 0, len(mp))
	for k := range mp {
		r = append(r, k)
	}
	return r
}

type ref struct {
	page     string
	pageName string
	heading  string
	ref      string
}

func (r ref) link() string {
	if r.ref != "" {
		return r.page + "#" + r.ref
	}
	return r.page

}

func (r ref) MarshalJSON() (bts []byte, err error) {
	x := struct {
		Page    string
		Heading string
		Link    string
	}{
		Page:    r.pageName,
		Heading: r.heading,
		Link:    r.link(),
	}
	bts, err = json.Marshal(x)
	return
}

func (r ref) depth() int {
	return strings.Count(r.page, `/`)
}

type cracker struct {
	debug   bool
	links   map[string][]ref
	pth     string
	page    string
	heading string
	pref    string
}

// extractLinks will walk a markdown document and attempt to resolve items from it to insert links
func (c *cracker) extractLinks(root ast.Node) (err error) {
	return c.walkNode(root, "")
}

func (c *cracker) ref() (r ref) {
	r = ref{
		page:     c.pth,
		pageName: c.page,
		heading:  c.heading,
		ref:      c.pref,
	}
	return
}

func (c *cracker) walkNode(n ast.Node, indent string) error {
	if c.debug && false {
		if indent == "" {
			fmt.Printf("%s\n", c.pth)
		} else {
			fmt.Printf("%s %s [%s]\n", indent, val(n), c.ref())
		}
	}
	if v, ok := n.(*ast.Heading); ok {
		if v != nil && len(v.Children) == 1 {
			if t, ok := v.Children[0].(*ast.Text); ok && t.Literal != nil {
				//this is a heading, so update our paragraph reference
				if c.page == `` {
					c.page = string(t.Literal)
				} else {
					c.heading = string(t.Literal)
					c.pref = getHeadingAnchor(t.Literal)
				}
				return c.processNode(t)
			}
		}
	}
	if err := c.processNode(n); err != nil {
		return err
	}
	for _, ch := range n.GetChildren() {
		if err := c.walkNode(ch, indent+" "); err != nil {
			return err
		}
	}
	return nil
}

// processNode determines whether or not we are going to "index" the given markdown node
func (c *cracker) processNode(n ast.Node) error {
	var bts []byte
	switch v := n.(type) {
	case *ast.Paragraph:
		bts = v.Literal
	case *ast.Text:
		bts = v.Literal
	case *ast.Code:
		bts = v.Literal
	case *ast.CodeBlock:
		bts = v.Literal
	case *ast.Heading:
		bts = v.Literal
	case *ast.TableCell:
		bts = v.Literal
	case *ast.ListItem:
		bts = v.Literal
	case *ast.Link:
		bts = v.Title
	case *ast.Image:
		bts = v.Title
	}
	if len(bts) == 0 {
		return nil
	}
	return c.processString(strings.ToLower(string(bts)))
}

func (c *cracker) processString(bts string) error {
	scanner := bufio.NewScanner(strings.NewReader(bts))
	// Set the split function for the scanning operation.
	scanner.Split(scanwords)
	for scanner.Scan() {
		if txt, ok := indexableString(scanner.Text()); ok {
			if strings.ContainsAny(txt, " \t\n") {
				//process the sub strings
				return c.processString(txt)
			}
			c.addRef(txt)
		}
	}
	if err := scanner.Err(); err != nil {
		if c.debug {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
	return nil
}

func (c *cracker) addRef(txt string) {
	if len(txt) == 0 {
		return
	}
	lnk := c.ref()
	if v, ok := c.links[txt]; ok {
		if !inRefList(lnk, v) {
			v = append(v, c.ref())
			c.links[txt] = v
		}
	} else {
		c.links[txt] = []ref{lnk}
	}
	if c.debug {
		fmt.Printf("%q -> %s\n", txt, lnk)
	}
}

func indexableString(txt string) (r string, ok bool) {
	if len(txt) == 0 {
		return
	}
	if r = strings.Trim(txt, "\"\n\t"); len(r) == 0 {
		return
	}

	//check if it only contains special characters or isn't printable
	for _, rn := range r {
		if !unicode.IsPrint(rn) {
			return
		}
		//check that SOMETHING is a non-symbol
		if unicode.IsLetter(rn) || unicode.IsNumber(rn) {
			ok = true
			break
		}
	}
	return
}

func getHeadingAnchor(bts []byte) string {
	mf := func(x rune) (r rune) {
		switch x {
		case ' ', '\t', '\n':
			r = '_'
		default:
			r = x
		}
		return
	}
	return string(bytes.Map(mf, bts))
}

func val(n ast.Node) string {
	switch v := n.(type) {
	case *ast.Paragraph:
		return fmt.Sprintf("Text: %s", string(v.Literal))
	case *ast.Text:
		return fmt.Sprintf("Text: %s", string(v.Literal))
	case *ast.Code:
		return fmt.Sprintf("Code: %s", string(v.Literal))
	case *ast.CodeBlock:
		return fmt.Sprintf("CodeBlock: %s", string(v.Literal))
	case *ast.Heading:
		if len(v.Children) == 1 && v.Children[0] != nil {
			if c, ok := v.Children[0].(*ast.Text); ok && len(c.Literal) != 0 {
				return fmt.Sprintf("Heading: %s", string(c.Literal))
			}
		}
		return fmt.Sprintf("Heading: %s", v.HeadingID)
	}
	return fmt.Sprintf("%T %+v", n, n)
}

func isSplit(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsControl(r) || !unicode.IsPrint(r) || unicode.IsSymbol(r)
}

func scanwords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSplit(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isSplit(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func inStringList(v string, sl []string) bool {
	for i := range sl {
		if sl[i] == v {
			return true
		}
	}
	return false
}

func inRefList(v ref, sl []ref) bool {
	for i := range sl {
		if sl[i] == v {
			return true
		}
	}
	return false
}

// sortLinks sorts a set of reference links by the following structure
// whether this is a search module
// most shallow (e.g. you aren't deep into the system"
// page
// href
func sortLinks(r []ref) {
	sort.SliceStable(r, func(i, j int) bool {
		//changelogs are always at the end
		icl := isChangeLogRef(r[i].page)
		jcl := isChangeLogRef(r[j].page)
		if icl != jcl {
			return !icl
		}
		//search modules are always up front
		isr := isSearchRef(r[i].page)
		jsr := isSearchRef(r[j].page)
		if isr != jsr {
			return isr
		}
		id := r[i].depth()
		jd := r[j].depth()
		if id < jd {
			return true //less
		} else if jd < id {
			return false //greater
		}
		//same depth, check the page
		if r[i].pageName == r[j].pageName {
			return r[i].ref < r[j].ref
		}

		return r[i].pageName < r[j].pageName
	})
}

func isSearchRef(pg string) bool {
	return strings.Contains(pg, `search/`)
}

func isChangeLogRef(pg string) bool {
	return strings.Contains(pg, `changelog/`)
}

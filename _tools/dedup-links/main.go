package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"

	"golang.org/x/net/html"

	mapset "github.com/deckarep/golang-set/v2"
)

// Parsed CLI arguments
var args struct {
	BaseUrl string `help:"base URL to use when constructing absolute paths" default:"/"`
	Workers int    `default:"4"`

	Root string `arg:"positional,required" help:"root directory to walk when replacing links"`
}

// Describes an "original" asset. An asset that will be used in places of duplicates.
type Replacement struct {
	// The link to use in HTML. An absolute path rooted at BaseURL
	link string

	// The absolute path to the asset on disk
	path string
}

func main() {
	arg.MustParse(&args)

	if s, err := os.Stat(args.Root); err != nil {
		log.Fatal(err)
	} else if !s.IsDir() {
		log.Fatal("root must be a directory")
	}

	// Walk files starting at Root, computing the hash for each file, creating hash => filepath map
	files := sync.Map{}
	computeFileHashes(args.Root, &files)

	// Compute links and construct hash => Replacement map
	replacements := sync.Map{}
	files.Range(func(key, value any) bool {
		hash := key.(string)
		path := value.(string)

		if filepath.Ext(path) == ".html" {
			return true
		}

		absLink, err := filepath.Rel(args.Root, path)
		if err != nil {
			log.Fatal(err)
		}

		absLink = filepath.Join(args.BaseUrl, absLink)

		replacements.Store(hash, Replacement{link: absLink, path: path})

		return true
	})

	// Walk files stargting at Root, using Replacement map to replace links
	// and accumulate set of replaced files
	oldFiles := replaceLinks(args.Root, &replacements)

	// Remove replaced files
	for _, path := range mapset.Sorted(oldFiles) {
		if path == "" {
			continue
		}
		os.Remove(path)
	}
}

// Walks root, computing the sha256 sum of each file, and updating hashMap with hash => path pairs.
// If a hash is already in the map, the value for the hash will not be updated.
func computeFileHashes(root string, hashMap *sync.Map) {
	pathChan := walkFiles(root)

	wg := sync.WaitGroup{}
	for i := 0; i < args.Workers; i++ {
		wg.Add(1)
		go func() {
			for path := range pathChan {
				hash := computeFileHash(path)
				hashMap.LoadOrStore(hash, path) // only store if it's not already in there
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// Computes the sha256 sum of the given file path
func computeFileHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Walks HTML files, starting at root, modifying files to replace links with those in Replacements
// Returns a Set of file paths that are no longer referenced
func replaceLinks(root string, replacements *sync.Map) mapset.Set[string] {
	pathChan := walkFiles(root)

	toDelete := mapset.NewSet[string]()

	wg := sync.WaitGroup{}
	for i := 0; i < args.Workers; i++ {
		wg.Add(1)
		go func() {
			for path := range pathChan {
				if filepath.Ext(path) != ".html" {
					continue
				}
				replaceLinksInFile(path, replacements, toDelete)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return toDelete
}

// Modifies the file at path by applying link replacements.
// Updates toDelete by adding the paths of files that were replaced
func replaceLinksInFile(path string, replacements *sync.Map, toDelete mapset.Set[string]) {
	s, err := os.Stat(args.Root)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(path, os.O_RDWR, s.Mode().Perm())
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	modified := replaceLinkFromNode(doc, path, replacements, toDelete)
	if !modified {
		return
	}

	err = f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	html.Render(f, doc)
	if err != nil {
		log.Fatal(err)
	}

}

// Given the path of the current file and src URL link to a relative file, returns the
// link that should replace srcUrl.
// If the replacement for srcUrl points to a different file than it did before, then "oldFile" is the
// path of the file that is no longer referenced
func lookupNewLink(thisFile, srcUrl string, newLinks *sync.Map) (newLink string, oldFile string) {

	// linkedFile is an absolute path
	linkedFile := ""
	if filepath.IsAbs(srcUrl) {
		linkedFile = filepath.Clean(filepath.Clean(args.Root) + filepath.Clean(srcUrl))
	} else {
		linkedFile = filepath.Join(filepath.Dir(thisFile), srcUrl)
	}

	hash := computeFileHash(linkedFile)

	linkInfoA, found := newLinks.Load(hash)
	if !found {
		log.Fatalln("Bad link")
	}

	linkInfo := linkInfoA.(Replacement)

	newLink = linkInfo.link

	if linkInfo.path != linkedFile {
		// if the new link references a different file, then we return the file referenced by the old link
		oldFile = linkedFile
	}

	return
}

// Walks an HTML document rooted at n, replacing links as required
// Updates toDelete with those files that are no longer referenced by links
// Returns true if the document was modified, otherwise false
func replaceLinkFromNode(n *html.Node, thisFile string, newLinks *sync.Map, toDelete mapset.Set[string]) (modified bool) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for i, a := range n.Attr {
			if a.Key == "src" {
				newLink, oldFile := lookupNewLink(thisFile, a.Val, newLinks)
				n.Attr[i].Val = newLink
				toDelete.Add(oldFile)
				modified = true
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childModified := replaceLinkFromNode(c, thisFile, newLinks, toDelete)
		modified = childModified || modified
	}

	return modified
}

// Given a file path root, recursively walks the directory, emitting paths over the returned channel
func walkFiles(root string) chan string {
	pathChan := make(chan string)

	go func() {
		filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if d.Type().IsRegular() {
				pathChan <- path
			}
			return nil
		})

		close(pathChan)
	}()

	return pathChan
}

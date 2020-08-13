const goToSearch = function(form) {
    const kwd = form.querySelector('input[type="search"]').value;
    console.log('go')
    window.location.href = '/?#!search-results.md?q=' + encodeURI(kwd);
}


const populate = function (links, kwd) {
    const container = document.querySelector('#search-results');
    if (!Array.isArray(links))  {
        links = [];
    }

    if (links.length === 0) {
        container.innerHTML = 'No results found.';
        return;
    }

    // api comes with results ordered by page
    let html = '<p>' + links.length + ' results found.</p>';
    let previousPage;
    links.forEach(l => {
        if (l.Page !== previousPage) {
            html += '<h2>' + l.Page + '</h2>';
        }
        if (!l.Heading || l.Heading.trim().length === 0) {
            l.Heading = l.Page;
        }
        html += `<p><a href="${l.Link}?q=${kwd}">
           ${l.Heading}
            </a></p>`
        previousPage = l.Page;
    });

    const el = document.createElement('div');
    el.innerHTML = html;
    container.appendChild(el);
}

const search = function(form) {
    const kwd = addSearchKeyword();
    if (kwd.length === 0) {
        populate([], kwd);
        return;
    }

    const url = '/api/search';
    $.ajax({
        url: url,
        type: 'POST',
        data: JSON.stringify({value: kwd}),
        success: function(resp) {
            if (resp && Array.isArray(resp.links)) {
                populate(resp.links, kwd);
            } else {
                populate([], kwd);
            }
        },
        dataType: 'json',
        contentType:"application/json",
    });
}

const addSearchKeyword = function() {
    const split = window.location.href.split('?q=');
    let kwd = '';
    if (split.length === 2) {
        kwd = decodeURIComponent(split[1]).trim();
    }
    fillSearch(kwd);
    return kwd;
}

const fillSearch = function(kwd) {
    const field = document.querySelector('#search-field');
    if (field) {
        field.value = kwd;
    } else {
        setTimeout(() => fillSearch(kwd), 100)
    }

}

$(document).ready(function() {
    $.ajax({
        url: '/api/search',
        type: 'HEAD',
        success: function() {
            const field = document.querySelector('#search-field');
            field.style.display = null;
        },
    });
});

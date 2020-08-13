const SESS_Q = 'q';

const goToSearch = function(form) {
    const kwd = form.querySelector('input[type="search"]').value;
    window.location.href = '/?#!search-results.md?q=' + encodeURI(kwd);
    window.sessionStorage.setItem(SESS_Q, kwd);
}


const populate = function (links) {
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
        html += `<p><a href="${l.Link}">
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
    console.log(kwd)
    if (!kwd || kwd.length === 0) {
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
                populate(resp.links);
            } else {
                populate([]);
            }
        },
        dataType: 'json',
        contentType:"application/json",
    });
}

const addSearchKeyword = function() {
    let kwd = window.sessionStorage.getItem(SESS_Q);
    if (kwd) {
        kwd = kwd.trim();
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

// #region CONSTANTS
const HEADERS_JSON = {
    'content-type': 'application/json;charset=UTF-8'
}

const HEADERS_HTML = {
    'content-type': 'text/html;charset=UTF-8'
}

const LINKS = {
    API_LINKS: [

        {
            "name": "Neighborhood Knockout (Global Game Jam 2020)",
            "url": "https://globalgamejam.org/2020/games/neighborhood-knockout-6"
        },
        {
            "name": "Octo Profile",
            "url": "https://octoprofile.now.sh/user?id=LeoLe101"
        },
        {
            "name": "Gams For Love",
            "url": "https://gamesforlove.org/"
        },
        {
            "name": "Magic Run For Fun",
            "url": "https://leole101.github.io/MagicRun.github.io/"
        }
    ],

    SOCIAL_LINKS: [
        {
            "url": "https://www.linkedin.com/in/leole101/",
            "svg": "https://simpleicons.org/icons/linkedin.svg"
        },
        {
            "url": "https://github.com/LeoLe101",
            "svg": "https://simpleicons.org/icons/github.svg"
        }
    ]
}

const DEBUG = true;

const REQUEST_LINK = "https://static-links-page.signalnerve.workers.dev";
// #endregion

// #region HTML REWRITER Helper
class LinksTransformer {
    constructor(links) { this.links = links; }

    async element(element) {
        this.links.forEach(link => {
            element.append(`<a href="${link.url}">${link.name}</a>`, { html: true });
        });
    }
}

class SocialLinksTransformer {
    constructor(links) { this.links = links; }

    async element(element) {
        element.removeAttribute('style');
        this.links.forEach(link => {
            element.append(
                `<a href="${link.url}">
                    <svg>
                        <img src="${link.svg}"/>
                    </svg>
                </a>`,
                { html: true });
        });
    }
}

const TitleTransformer = {
    element: el => {
        el.setInnerContent('Leo Le - Lots of links about Leo!');
    }
}

const BackgroundTransformer = {
    element: el => {
        el.setAttribute('style', 'background-image: linear-gradient(to bottom right, #D002FE)')
    }
}

const ProfileTransformer = {
    element: el => {
        el.removeAttribute('style');
    }
}

const UserNameTransformer = {
    element: el => {
        el.setInnerContent('Leo Le');
    }
}

const ImageTransformer = {
    element: el => {
        el.setAttribute('src', 'https://avatars2.githubusercontent.com/u/22701345?v=4');
    }
}

const Rewriter = new HTMLRewriter()
    .on('div#profile', ProfileTransformer)
    .on('img#avatar', ImageTransformer)
    .on('h1#name', UserNameTransformer)
    .on('div#links', new LinksTransformer(LINKS.API_LINKS))
    .on('div#social', new SocialLinksTransformer(LINKS.SOCIAL_LINKS))
    .on('title', TitleTransformer)
    .on('body', BackgroundTransformer)
// #endregion

// #region API Handler

addEventListener('fetch', event => {
    try {
        event.respondWith(handleRequest(event.request))
    } catch (e) {
        // Incase server error
        if (DEBUG) {
            return event.respondWith(
                new Response(e.message || e.toString(), {
                    status: 500,
                }),
            )
        }
        event.respondWith(new Response('Internal Error', { status: 500 }))
    }
})

/**
 * Respond to all requests
 *  - Check if the path is '/links' to return JSON Response
 *  - Else, return the transformed static HTML
 *  
 * @param {Request} request coming request from client's browser
 */
async function handleRequest(request) {
    let response;
    let url = new URL(request.url);
    let urlPath = url.pathname;

    if (urlPath === '/links') {
        // Return API Links as JSON
        response = new Response(JSON.stringify(LINKS.API_LINKS, null, 2), HEADERS_JSON);
    } else {
        let staticHTMLResponse = await fetch(REQUEST_LINK);
        // Transform HTML when requested successfully
        if (staticHTMLResponse.ok) {
            response = Rewriter.transform(staticHTMLResponse);
            response.headers = HEADERS_HTML;
        } else {
            response = new Response('Static HTML Request Error', { status: 500 });
        }
    }
    return response;
}

// #endregion


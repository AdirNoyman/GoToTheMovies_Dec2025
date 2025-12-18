
export class YouTubeEmbed extends HTMLElement {


    static get observedAttributes() {
        return ['data-url']
    }


    attributeChangedCallback(prop, value) {

        if (prop === 'data-url') {
            const url = this.dataset.url
            console.log(url)

            let urlID = url.toString().split('=')
            urlID = urlID[urlID.length - 1]
            console.log("UrlID 👉", urlID )
            let test = `https://www.youtube.com/embed/${urlID}` 
            console.log("Test:", test)

            this.innerHTML = `<iframe width="100%" height="300" src="https://www.youtube.com/embed/${urlID}" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>`
        }
    }

}

customElements.define("youtube-embed", YouTubeEmbed)
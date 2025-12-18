export class MovieItem extends HTMLElement {
  constructor(movie) {
    super();
    this.movie = movie;
  }

  // Lifecycle method called when the element is added (mounted) to the DOM
  connectedCallback() {
    this.innerHTML = `
               <a href="#">
               <article>
                 <img src="${this.movie.poster_url}" alt="${this.movie.title} poster" />
                 <p>${this.movie.title} (${this.movie.release_year})</p>                
               </article>
               </a>
            `;
  }
}

// Register (inject...) the custom element for use by the DOM HTML
customElements.define('movie-item', MovieItem);

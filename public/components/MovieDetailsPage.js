import { API } from '../services/API.js';

export class MovieDetailsPage extends HTMLElement {
  movieId = null;
  movie = null;

  async render() {
    try {
      this.movie = await API.getMovieById(this.movieId);
    } catch (error) {
      // TODO: handle error
    }
    if (!this.movie) {
      this.textContent = 'Movie not found 😩';
      return;
    } else {
      const template = document.getElementById('template-movie-details');
      // cloneNode = true => make a deep copy of the template content
      const content = template.content.cloneNode(true);
      this.appendChild(content);

      // Title and tagline
      this.querySelector('h2').textContent = this.movie.title;
      this.querySelector('h3').textContent = this.movie.tagline;

      // Poster image
      const posterImg = this.querySelector('header img');
      posterImg.src = this.movie.poster_url;
      posterImg.alt = `${this.movie.title} Poster`;

      // Trailer
      const trailerEmbed = this.querySelector('#trailer');
      trailerEmbed.setAttribute('data-url', this.movie.trailer_url);

      // Movie overview
      this.querySelector('#overview').textContent = this.movie.overview;

      // Metadata
      const metadata = this.querySelector('#metadata');
      metadata.innerHTML = `
        <dt>Release Year</dt>
        <dd>${this.movie.release_year}</dd>
        <dt>Rating</dt>
        <dd>${this.movie.score}/10</dd>
        <dt>Popularity</dt>
        <dd>${this.movie.popularity.toFixed(1)}</dd>
        <dt>Language</dt>
        <dd>${this.movie.language.toUpperCase()}</dd>
      `;

      // Genres
      const genresList = this.querySelector('#genres');
      if (this.movie.genres && this.movie.genres.length > 0) {
        genresList.innerHTML = this.movie.genres
          .map((genre) => `<li>${genre.name}</li>`)
          .join('');
      } else {
        genresList.style.display = 'none';
      }

      // Overview
      this.querySelector('#overview').textContent = this.movie.overview;

      // Cast
      const castList = this.querySelector('#cast');
      if (this.movie.casting && this.movie.casting.length > 0) {
        castList.innerHTML = this.movie.casting
          .map(
            (actor) => `<li>
          <img src="${
            actor.image_url ?? '/images/generic_actor.jpg'
          }" alt="Picture of ${actor.first_name} ${actor.last_name}" />
          <p>${actor.first_name} ${actor.last_name}</p></li>`
          )
          .join('');
      } else {
        castList.style.display = 'none';
      }
    }
  }

  connectedCallback() {
    this.movieId = 140; // Hardcoded movie ID for demonstration
    this.render();
  }
}

// Register (inject...) the custom element for use by the DOM HTML
customElements.define('movie-details-page', MovieDetailsPage);

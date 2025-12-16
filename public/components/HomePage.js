import { API } from "../services/API.js";
import { MovieItem } from "./MovieItem.js";

export class HomePage extends HTMLElement {

   async render() {
        const topMovies = await API.getTopMovies();
        createMoviesCards(topMovies, document.getElementById('top-movies'));

        const randomMovies = await API.getRandomMovies();
        createMoviesCards(randomMovies, document.getElementById('random-movies'));

        function createMoviesCards(movies, ul) {
            ul.innerHTML = '';
            movies.forEach(movie => {
                const li = document.createElement('li');
                li.appendChild(new MovieItem(movie));
                ul.appendChild(li);
            });

        }

    }



    connectedCallback() {
        const template = document.getElementById('template-home');
        const content = template.content.cloneNode(true);
        this.appendChild(content);

        this.render();
    }
}

customElements.define('home-page', HomePage);
import { API } from './services/API.js';
// import './components/HomePage.js';
import './components/AnimatedLoading.js';
import './components/MovieDetailsPage.js';
import './components/YouTubeEmbed.js';

window.app = {
  search: (event) => {
    event.preventDefault();
    const query = document.getElementById('search-input').value;
  },

  api: API,
};

import { HomePage } from '../components/HomePage.js';
import { MovieDetailsPage } from '../components/MovieDetailsPage.js';
import { MoviePage } from '../components/MoviePage.js';

export const routes = [
  // Get all movies
  {
    path: '/',
    component: HomePage,
  },
  // Get movie by ID
  {
    // use regex to identifying the path
    path: /\/movies\/(\d+)/,
    component: MovieDetailsPage,
  },
  // Search movie
  {
    path: '/movies',
    component: MoviePage,
  },
];

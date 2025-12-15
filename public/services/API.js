export const API = {
  baseURL: '/api/movies',

  fetch: async (route, args) => {
    try {
      const queryParams = args ? new URLSearchParams(args).toString() : "";
      const response = await fetch(`${API.baseURL}/${route}`, queryParams ? `?${queryParams}` : "");
      const result = await response.json();
      return result;
    } catch (error) {
      console.error(`Error trying to fetch ${route}`, error);
      return null;
    }
  },

  getTopMovies: () => {
    return API.fetch('top');
  },
  getRandomMovies: () => {
    return API.fetch('random');
  },
  getMovieById: (id) => {
    return API.fetch(`movie/${id}`);
  },
  searchMovies: (q,order,genre) => {
    return API.fetch('search', {q,order,genre});
  }
};

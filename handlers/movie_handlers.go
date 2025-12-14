package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"	
	"devvvine.com/moviez/data"
	"devvvine.com/moviez/logger"
	"devvvine.com/moviez/models"
)

var movies = []models.Movie{
	{
		ID:          1,
		TMDB_ID:     278,
		Title:       "The Shawshank Redemption",
		ReleaseYear: 1994,
		Genres: []models.Genre{
			{ID: 1, Name: "Drama"},
			{ID: 2, Name: "Crime"},
		},
		Keywords: []string{"prison", "friendship", "hope", "redemption", "escape", "wrongful conviction"},
		Casting: []models.Actor{
			{
				ID:        1,
				FirstName: "Tim",
				LastName:  "Robbins",
			},
			{
				ID:        2,
				FirstName: "Morgan",
				LastName:  "Freeman",
			},
			{
				ID:        3,
				FirstName: "Bob",
				LastName:  "Gunton",
			},
		},
	},
	{
		ID:          2,
		TMDB_ID:     238,
		Title:       "The Godfather",
		ReleaseYear: 1972,
		Genres: []models.Genre{
			{ID: 1, Name: "Drama"},
			{ID: 2, Name: "Crime"},
		},
		Keywords: []string{"mafia", "family", "crime family", "sicily", "1940s", "organized crime"},
		Casting: []models.Actor{
			{
				ID:        4,
				FirstName: "Marlon",
				LastName:  "Brando",
			},
			{
				ID:        5,
				FirstName: "Al",
				LastName:  "Pacino",
			},
			{
				ID:        6,
				FirstName: "James",
				LastName:  "Caan",
			},
		},
	},
	{
		ID:          3,
		TMDB_ID:     155,
		Title:       "The Dark Knight",
		ReleaseYear: 2008,
		Genres: []models.Genre{
			{ID: 3, Name: "Action"},
			{ID: 2, Name: "Crime"},
			{ID: 1, Name: "Drama"},
		},
		Keywords: []string{"dc comics", "crime fighter", "joker", "superhero", "sequel", "vigilante"},
		Casting: []models.Actor{
			{
				ID:        7,
				FirstName: "Christian",
				LastName:  "Bale",
			},
			{
				ID:        8,
				FirstName: "Heath",
				LastName:  "Ledger",
			},
			{
				ID:        9,
				FirstName: "Aaron",
				LastName:  "Eckhart",
			},
		},
	},
	{
		ID:          4,
		TMDB_ID:     27205,
		Title:       "Inception",
		ReleaseYear: 2010,
		Genres: []models.Genre{
			{ID: 3, Name: "Action"},
			{ID: 4, Name: "Science Fiction"},
			{ID: 5, Name: "Thriller"},
		},
		Keywords: []string{"dream", "mission", "architecture", "subconscious", "heist", "mindbender"},
		Casting: []models.Actor{
			{
				ID:        10,
				FirstName: "Leonardo",
				LastName:  "DiCaprio",
			},
			{
				ID:        11,
				FirstName: "Joseph",
				LastName:  "Gordon-Levitt",
			},
			{
				ID:        12,
				FirstName: "Elliot",
				LastName:  "Page",
			},
		},
	},
	{
		ID:          5,
		TMDB_ID:     680,
		Title:       "Pulp Fiction",
		ReleaseYear: 1994,
		Genres: []models.Genre{
			{ID: 5, Name: "Thriller"},
			{ID: 2, Name: "Crime"},
		},
		Keywords: []string{"drug dealer", "boxer", "crime boss", "hitman", "briefcase", "nonlinear timeline"},
		Casting: []models.Actor{
			{
				ID:        13,
				FirstName: "John",
				LastName:  "Travolta",
			},
			{
				ID:        14,
				FirstName: "Samuel",
				LastName:  "Jackson",
			},
			{
				ID:        15,
				FirstName: "Uma",
				LastName:  "Thurman",
			},
		},
	},
}

type MovieHandler struct {
	Storage data.MovieStorage
	Logger  *logger.Logger
}

func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data any) error {

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		h.Logger.Error("Failed to encode JSON response", err)
		http.Error(w, "Failed to send movies due to internal server error 😫", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.Logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MovieHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

// Get all movies - currently returns hardcoded data
func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {

	movies, err := h.Storage.GetTopMoviesFromRepo()
	if err != nil {
		h.Logger.Error("Failed to get top movies from storage", err)
		http.Error(w, "Failed to get top movies due to internal server error 😫", http.StatusInternalServerError)
		return
	}

	h.writeJSONResponse(w, movies)

}

// Get random movie
// func (h *MovieHandler) GetRandomMovie(w http.ResponseWriter, r *http.Request) {

// 	movies, err := h.Storage.GetRandomMovies()
// 	if err != nil {
// 		h.Logger.Error("Failed to get random movies from storage", err)
// 		http.Error(w, "Failed to get random movies due to internal server error 😫", http.StatusInternalServerError)
// 		return
// 	}

// 	  // Prevent caching
//     w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
//     w.Header().Set("Pragma", "no-cache")
//     w.Header().Set("Expires", "0")

// 	// Get random movie from the hardcoded list
// 	randomMovie := movies[rand.IntN(len(movies))]

// 	h.writeJSONResponse(w, randomMovie)

// }

// Get random movies
func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {

	// TODO: fix logic to return multiple random movies without duplicates

	movies, err := h.Storage.GetRandomMoviesFromRepo()
	if err != nil {
		h.Logger.Error("Failed to get random movies from storage", err)
		http.Error(w, "Failed to get random movies due to internal server error 😫", http.StatusInternalServerError)
		return
	}

	h.writeJSONResponse(w, movies)

}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := h.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = h.Storage.SearchMoviesByNameFromRepo(query, order, genre)
	}
	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served movies 🤓🤘")
	}
}

func (h *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL path
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := h.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := h.Storage.GetMovieByIDFromRepo(id)
	if h.handleStorageError(w, err, "Failed to get movie by ID") {
		return
	}
	if h.writeJSONResponse(w, movie) == nil {
		h.Logger.Info("Successfully served movie with ID: " + idStr)
	}
}

func (h *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.Storage.GetAllGenresFromRepo()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	if h.writeJSONResponse(w, genres) == nil {
		h.Logger.Info("Successfully served genres")
	}
}

func NewMovieHandler(storage data.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		Storage: storage,
		Logger:  log,
	}
}
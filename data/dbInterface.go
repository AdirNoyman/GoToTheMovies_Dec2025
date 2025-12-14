package data

import "devvvine.com/moviez/models"

type MovieStorage interface {
	
	GetTopMoviesFromRepo() ([]models.Movie, error)
	GetRandomMoviesFromRepo() ([]models.Movie, error)
	GetMovieByIDFromRepo(id int) (models.Movie, error)
	SearchMoviesByNameFromRepo(name string, order string, genre *int) ([]models.Movie, error)
	GetAllGenresFromRepo() ([]models.Genre, error)
}
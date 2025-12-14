package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"devvvine.com/moviez/data"
	"devvvine.com/moviez/handlers"
	"devvvine.com/moviez/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func initializeLogger() *logger.Logger {

	logInstance, err := logger.NewLogger("./app.logs")
	if err != nil {
		log.Fatalf("Failed to initialize logger 😩👉: %v", err)
	}
	return logInstance
}

func main() {

	logInstance := initializeLogger()
	defer logInstance.Close() // defer ensures Close() is called when main() exits, even if program crashes or exits early
	// Environment variables
	err := godotenv.Load(".env")
	if err != nil {
		logInstance.Error("Error loading .env file", err)
	}

	// Connect to the database
	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		logInstance.Error("POSTGRES_URL is not set in environment variables", nil)
		log.Fatal("POSTGRES_URL is not set in environment variables")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logInstance.Error("Failed to connect to database 😫", err)
		log.Fatalf("Failed to connect to database 😫: %v", err)
	}

	// Close the database connection when main function exits
	fmt.Println("Database connected successfully 🤓🤘")
	defer db.Close()

	// Initialize data repository of movies
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		logInstance.Error("Failed to initialize movie repository 😫", err)
		log.Fatal("Failed to initialize movie repository 😫", err)
	}

	// Initialize the movie handler
	movieHandler := handlers.MovieHandler{}
	movieHandler.Storage = movieRepo
	movieHandler.Logger = logInstance

	// Serve static files from the "public" directory
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie) // expects ID in the path
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.Handle("/", http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")

	println("Server is running 🚀🤘🤓 on port", port)

	err = http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		logInstance.Error("Server failed to start 😫", err)
	}

}

////////////////////////////////////////////////////////////

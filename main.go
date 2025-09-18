package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"frontendmasters.com/reelingit/data"
	"frontendmasters.com/reelingit/handlers"
	"frontendmasters.com/reelingit/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	logInstance := initializeLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL is not set in the environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to create movie repository: %v", err)
	}

	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)

	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	catchAllHandler := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/index.html")
	}
	http.HandleFunc("/movies", catchAllHandler)
	http.HandleFunc("/movies/", catchAllHandler)
	http.HandleFunc("/account/", catchAllHandler)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}

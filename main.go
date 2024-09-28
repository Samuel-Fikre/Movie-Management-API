package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// The struct tags (e.g., json:"id") ensure that the fields are serialized to JSON with those specific names

//The pointer allows the Director field to be nil, meaning the movie may or may not have a director. This is useful if you want to represent the absence of a director for certain movies.

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"iSbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// Handler functions

// to set the HTTP response header to indicate that the response content type is JSON.

// Header() returns an http.Header object, which represents the response headers.
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w) creates a new JSON encoder that writes to the w object,
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(100000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//mux.Vars(r) retrieves the route variables (URL parameters) from the incoming HTTP request r. These variables are typically defined in the URL path as placeholders, like /movies/{id}.
	//mux.Vars(r) will return a map with the URL parameters. For example, if the URL is /movies/123, the map would be {"id": "123"}
	params := mux.Vars(r)
	//This Go code snippet is part of a typical operation where you're trying to remove an item from a slice (in this case, a list of movies) based on a condition, specifically when the ID of the item matches a certain parameter (likely extracted from the URL using something like mux.Vars(r)).
	for index, item := range movies {
		// In the code, append is like your glue. It’s combining (or gluing) the part of the list before the movie you want to remove, with the part after it. So just like you skipped over the Blue toy car, the code skips over the movie you want to remove and makes a new list without it!
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		ISBN:     "438277",
		Title:    "Movie One",
		Director: &Director{Firstname: "Christopher", Lastname: "Nolan"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		ISBN:     "200123",
		Title:    "Movie Two",
		Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"},
	})

	// Define routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))

}

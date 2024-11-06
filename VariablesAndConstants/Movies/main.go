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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
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
	movie.ID = strconv.Itoa(rand.Intn(10000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123451",
		Title: "Movie Title 1",
		Director: &Director{
			Firstname: "Director First 1",
			Lastname:  "Director Last 1",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "123452",
		Title: "Movie Title 2",
		Director: &Director{
			Firstname: "Director First 2",
			Lastname:  "Director Last 2",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "123453",
		Title: "Movie Title 3",
		Director: &Director{
			Firstname: "Director First 3",
			Lastname:  "Director Last 3",
		},
	})
	movies = append(movies, Movie{
		ID:    "4",
		Isbn:  "123454",
		Title: "Movie Title 4",
		Director: &Director{
			Firstname: "Director First 4",
			Lastname:  "Director Last 4",
		},
	})
	movies = append(movies, Movie{
		ID:    "5",
		Isbn:  "123455",
		Title: "Movie Title 5",
		Director: &Director{
			Firstname: "Director First 5",
			Lastname:  "Director Last 5",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movie{id}", updateMovie).Methods("PUT")
	fmt.Printf("starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

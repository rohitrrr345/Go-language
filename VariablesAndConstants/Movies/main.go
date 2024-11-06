 package main
  import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
  )
  type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json :"director"`
  }
  type Director struct {
	firstname string `json:"firstname"`
	lastname string `json:"lastname"`
  }
  var movies []Movie
  func main(){
	r:=mux.NewRouter();
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123451",
		Title: "Movie Title 1",
		Director: &Director{
			firstname: "Director First 1",
			lastname:  "Director Last 1",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "123452",
		Title: "Movie Title 2",
		Director: &Director{
			firstname: "Director First 2",
			lastname:  "Director Last 2",
		},
	})
	movies = append(movies, Movie{
		ID:    "3",
		Isbn:  "123453",
		Title: "Movie Title 3",
		Director: &Director{
			firstname: "Director First 3",
			lastname:  "Director Last 3",
		},
	})
	movies = append(movies, Movie{
		ID:    "4",
		Isbn:  "123454",
		Title: "Movie Title 4",
		Director: &Director{
			firstname: "Director First 4",
			lastname:  "Director Last 4",
		},
	})
	movies = append(movies, Movie{
		ID:    "5",
		Isbn:  "123455",
		Title: "Movie Title 5",
		Director: &Director{
			firstname: "Director First 5",
			lastname:  "Director Last 5",
		},
	})

	r.handleFunc("/movies",getMovies).Methods("GET")
	r.handleFunc("/movies/{id}",getMovie).Methods("GET")
	r.handleFunc("/movies",createMovie).Methods("POST")
	r.handleFunc("/movie{id}",deleteMovie).Methods("DELETE")
	r.handleFunc("/movie{id}",updateMovie).Methods("PUT")




                   
  }                                       
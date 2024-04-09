package main

import (
	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "18953", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Malcovic"}})
	movies = append(movies, Movie{ID: "2", Isbn: "29045", Title: "Movie Two", Director: &Director{Firstname: "Will", Lastname: "Smith"}})

}

package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//main structure
type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string "json:firstname"
	Lastname string "json:firstname"
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params :=mux.Vars(r)
	for _,item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	//decoding for storing in pc
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies,movie)
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	//set movie content type
	w.Header().Set("Content-Type","application/json")
	//params
	params := mux.Vars(r)
	//loop over  the movies and gives us the one with matched id
	for index,item := range movies{
		if item.ID == params["id"]{
			//delete the  movie with the id sent by user
			movies = append(movies[:index],movies[index+1:]...)
	//add a new movie - the movie we send in the body of postman
			//decoding for storing in pc
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies,movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	
}


func main() {
	//init router
	r :=mux.NewRouter()

	// initial values
	movies = append(movies,Movie{
		ID:"1",
		Isbn:"12312",
		Title:"Movie One",
		Director:&Director{
			Firstname: "John",
			Lastname: "Wick",
		}})
	
	movies = append(movies,Movie{
		ID:"1",
		Isbn:"12312",
		Title:"Movie One",
		Director:&Director{
			Firstname: "John",
			Lastname: "Wick",
		}})

	//defining routers
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":8080",r))
}
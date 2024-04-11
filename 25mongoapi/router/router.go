package router

import (
	"github.com/Swaty-G/GolangTutorial/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	//Get all movies
	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	//Create a movie
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	//Update a movie to mark as watched
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	//Delete a movie
	router.HandleFunc("/api/movies/{id}", controller.DeleteAMovie).Methods("DELETE")
	//Delete all movies
	router.HandleFunc("/api/deleteAllMovie", controller.DeleteAllMovies).Methods("DELETE")

	return router

}

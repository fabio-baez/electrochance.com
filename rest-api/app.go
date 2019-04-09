package main

import (
	"log"
	"net/http"

	"electrochance.com/rest-api/config"
	"electrochance.com/rest-api/controllers"

	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//var conf = config.Config{}
var conf = config.YAML{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	//conf.Read()
	conf.ReadYaml()
}

// Define HTTP request routes
func main() {

	r := mux.NewRouter()

	//APIS
	r.HandleFunc("/api/v1.0/movies", controllers.AllMovies).Methods("GET")
	r.HandleFunc("/api/v1.0/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/api/v1.0/movies", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/api/v1.0/movies", controllers.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/v1.0/movies/{id}", controllers.FindMovie).Methods("GET")

	//r.Use(app.JwtAuthentication)

	port := conf.Port
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

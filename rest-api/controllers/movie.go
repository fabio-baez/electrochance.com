package controllers

import (
	"encoding/json"
	"net/http"

	"electrochance.com/rest-api/config"
	"electrochance.com/rest-api/models"
	"electrochance.com/rest-api/dao"

	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

var conf = config.YAML{}
var do = dao.MoviesDAO{}

func init() {

	conf.ReadYaml()

	do.Server = conf.Server
	do.Database = conf.Database

	do.Connect()
}

// AllMovies GET list of movies
func AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := do.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, movies)
}

// FindMovie GET a movie by its ID
func FindMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movie, err := do.FindByID(params["id"])
	
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJSON(w, http.StatusOK, movie)
}

// CreateMovie POST a new movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := do.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, movie)
}

// UpdateMovie PUT update an existing movie
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := do.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

// DeleteMovie DELETE an existing movie
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	
	defer r.Body.Close()
	var movie models.Movie

	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := do.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Team-Farr/go-farr-api/model"
	"github.com/gorilla/mux"
)

//CreateNews receives News and stores them and return the News IDs
func CreateNews(w http.ResponseWriter, r *http.Request) {
	var news *model.News

	if err := decodeAndValidate(r, news); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
		return
	}
	newsJSON, err := json.Marshal(news)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newsJSON)
}

//GetNews gets all news in a list with an optional news item specified by ID
func GetNews(w http.ResponseWriter, r *http.Request) {}

//OverwriteNews overwrites an existing news item with ID and overwrite it with a complete News object
func OverwriteNews(w http.ResponseWriter, r *http.Request) {}

//EditNews edits an existing news item with ID and replace it with an incomplete News object
func EditNews(w http.ResponseWriter, r *http.Request) {}

//DeleteNews deletes an existing news item with ID
func DeleteNews(w http.ResponseWriter, r *http.Request) {}

func decodeAndValidate(r *http.Request, v model.ApiValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	return v.Validate(r)
}

func main() {
	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/news", GetNews).Methods("GET")
	apiV1.HandleFunc("/news", CreateNews).Methods("POST")
	apiV1.HandleFunc("/news/{ID}", GetNews).Methods("GET")
	apiV1.HandleFunc("/news/{ID}", OverwriteNews).Methods("PUT")
	apiV1.HandleFunc("/news/{ID}", EditNews).Methods("PATCH")
	apiV1.HandleFunc("/news/{ID]", DeleteNews).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

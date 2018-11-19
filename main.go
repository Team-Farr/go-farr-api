package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GetNews returns all the news in a list format
func GetNews(w http.ResponseWriter, r *http.Request) {
}

//CreateNews receives News and stores them and return the NewsItem IDs
func CreateNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NEWS")
}

//GetNewsItem gets a news item specified by ID
func GetNewsItem(w http.ResponseWriter, r *http.Request) {}

//OverwriteNewsItem overwrites an existing news item with ID and overwrite it with a complete News object
func OverwriteNewsItem(w http.ResponseWriter, r *http.Request) {}

//EditNewsItem edits an existing news item with ID and replace it with an incomplete News object
func EditNewsItem(w http.ResponseWriter, r *http.Request) {}

//DeleteNewsItem deletes an existing news item with ID
func DeleteNewsItem(w http.ResponseWriter, r *http.Request) {}

func main() {
	router := mux.NewRouter()
	apiV1 := router.PathPrefix("/api/v1").Subrouter()
	apiV1.HandleFunc("/news", GetNews).Methods("GET")
	apiV1.HandleFunc("/news", CreateNews).Methods("POST")
	apiV1.HandleFunc("/news/{ID}", GetNewsItem).Methods("GET")
	apiV1.HandleFunc("/news/{ID}", OverwriteNewsItem).Methods("PUT")
	apiV1.HandleFunc("/news/{ID}", EditNewsItem).Methods("PATCH")
	apiV1.HandleFunc("/news/{ID]", DeleteNewsItem).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

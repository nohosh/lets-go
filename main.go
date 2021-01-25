package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepa API hit")
}

func allArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All articles with mux Router")
	articles := Articles{
		Article{Title: "Test Title", Desc: "Sample Desc", Content: "random contentt is typed here sos please ignore"},
	}
	json.NewEncoder(w).Encode(articles)
}

func samplePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post reuquest is hit")
}
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/articles", allArticle).Methods("GET")
	r.HandleFunc("/articles", samplePost).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", r))
}
func main() {
	handleRequests()
}

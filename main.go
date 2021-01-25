package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	fmt.Println("All articles")
	articles := Articles{
		Article{Title: "Test Title", Desc: "Sample Desc", Content: "random contentt is typed here sos please ignore"},
	}
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticle)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func main() {
	handleRequests()
}

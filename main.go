package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HTTP routes

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome To Server")
}

func addABook(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var book Book
	err := decoder.Decode(&book)
	if err != nil {
		panic(err)
	}
	AddBook(book.Name, book.Author, book.ReleaseDate, book.Pages)

}

func getAllBooks(w http.ResponseWriter, req *http.Request) {
	books := GetBooks()
	jsonType, _ := json.Marshal(books)
	fmt.Fprintf(w, string(jsonType))
}

func searchABook(w http.ResponseWriter, req *http.Request) {
	id, ok := req.URL.Query()["id"]
	if ok {
		i, _ := strconv.ParseInt(id[0], 10, 64)
		result := SearchBookByID(i)
		jsonForm, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonForm))
	}

	name, ok := req.URL.Query()["name"]
	if ok {
		result := SearchBookByName(name[0])
		jsonForm, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonForm))
	}

	author, ok := req.URL.Query()["author"]
	if ok {
		result := SearchBookByAuthor(author[0])
		jsonForm, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonForm))
	}

	release, ok := req.URL.Query()["release"]
	if ok {
		result := SearchBookByReleaseDate(release[0])
		jsonForm, _ := json.Marshal(result)
		fmt.Fprintf(w, string(jsonForm))
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", index)
	route.HandleFunc("/add", addABook)
	route.HandleFunc("/get", getAllBooks)
	route.HandleFunc("/search", searchABook)
	http.Handle("/", route)

	http.ListenAndServe(":5555", nil)
}

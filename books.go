package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

var jsonPath = "books.json"

// Books x
type Books struct {
	List []Book `json:"books"`
}

// Book x
type Book struct {
	ID          int64  `json:"Id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	ReleaseDate string `json:"release_date"`
	Pages       int    `json:"pages"`
}

// json funcs

func genID() int64 {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * 150)
	r := rand.Int63n(1e7)
	return r
}

// GetBooks x
func GetBooks() Books {
	reader, _ := ioutil.ReadFile(jsonPath)
	var books Books
	if err := json.Unmarshal(reader, &books); err != nil {
		panic(err)
	}
	return books
}

// AddBook x
func AddBook(name, author, releaseDate string, pages int) {
	books := GetBooks()
	newBook := Book{genID(), name, author, releaseDate, pages}
	books.List = append(books.List, newBook)
	newBooks, _ := json.Marshal(books)
	ioutil.WriteFile(jsonPath, newBooks, 0777)
}

// SearchBookByAuthor x
func SearchBookByAuthor(author string) []Book {
	books := GetBooks()
	var result []Book
	for _, v := range books.List {
		if v.Author == author {
			result = append(result, v)
		}
	}
	return result
}

// SearchBookByName x
func SearchBookByName(name string) []Book {
	books := GetBooks()
	var result []Book
	for _, v := range books.List {
		if v.Name == name {
			result = append(result, v)
		}
	}
	return result
}

// SearchBookByReleaseDate x
func SearchBookByReleaseDate(releaseDate string) []Book {
	books := GetBooks()
	var result []Book
	for _, v := range books.List {
		if v.ReleaseDate == releaseDate {
			result = append(result, v)
		}
	}
	return result
}

// SearchBookByID x
func SearchBookByID(ID int64) []Book {
	books := GetBooks()
	var result []Book
	for _, v := range books.List {
		if v.ID == ID {
			result = append(result, v)
		}
	}
	return result
}

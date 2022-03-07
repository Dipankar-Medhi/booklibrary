package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Dipankar-Medhi/go-booklibrary/pkg/models"
	"github.com/Dipankar-Medhi/go-booklibrary/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book // NewBook of type Book struct

//get book function
func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//get book by Id
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0) // convert into str
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails) //send a json response to user received from db
	w.Header().Set("Content-Type", "pkglication/josn")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}   //we received in json
	utils.ParseBody(r, createBook) // parsed it so that our db can understand
	b := createBook.CreateBook()
	res, _ := json.Marshal(b) // convert record to json to send to user using postman
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

//delete book function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID) // DeleteBook return the deleted book
	res, _ := json.Marshal(book)  // convert the book return from db into json
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

//update function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, dB := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	dB.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

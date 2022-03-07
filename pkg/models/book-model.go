package models

import (
	"github.com/Dipankar-Medhi/go-booklibrary/pkg/config"
	"github.com/jinzhu/gorm"
)

var dB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	config.Connect()
	dB = config.GetDB()
	dB.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	dB.NewRecord(b)
	dB.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	dB.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	dB := dB.Where("ID = ?", Id).Find(&getBook)
	return &getBook, dB
}

func DeleteBook(ID int64) Book {
	var book Book
	dB.Where("ID=?", ID).Delete(book)
	return book
}

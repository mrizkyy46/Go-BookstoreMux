package models

import (
	"github.com/mrizkyy46/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	ISBN        string `json:"ISBN"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getbook Book
	db := db.Where("id=?", id).Find(&getbook)
	return &getbook, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("id=?", id).Delete(&book)
	return book
}

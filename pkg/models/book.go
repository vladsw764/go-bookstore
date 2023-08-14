package models

import (
	"github.com/vladsw764/go-bookstore/pkg/configs"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	configs.Connect()
	db = configs.GetDb()
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return
	}
}

func (b *Book) CreateBook() *Book {
	if b.ID == 0 {
		db.Create(&b)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID = ?", ID).Find(&getBook)
	return &getBook, db
}

func DeleteBookByID(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}

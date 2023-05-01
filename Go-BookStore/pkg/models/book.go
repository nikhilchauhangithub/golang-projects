package models

import (
	"github.com/Go-BookStore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}

func init(){
	config.Connect()
	db=config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
db.Create(&b)
 return b
}

func GetAllBooks() []Book{
	var Books []Book
	db.Find(&Books)
	return Books
}
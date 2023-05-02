package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Go-BookStore/pkg/models"
	"github.com/Go-BookStore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id, err:=strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _:=models.GetBookById(Id)
	res,_:=json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request)  {
	CreateBook:=&models.Book{}
	utils.ParseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request)  {
	vars:=mux.Vars(r)
	bookId:=vars["bookId"]
	Id, err:=strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBook(Id)
	res,_:=json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook (w http.ResponseWriter, r *http.Request){
var updateBook =&models.Book{}
utils.ParseBody(r,updateBook)
vars:=mux.Vars(r)
bookId:=vars["bookId"]
Id,err:=strconv.ParseInt(bookId,0,0)
if err!= nil{
	fmt.Println("error while parsing")
}
bookDetails,db:=models.GetBookById(Id)
if updateBook.Name !=""{
	bookDetails.Name = updateBook.Name
}
if updateBook.Author !=""{
	bookDetails.Author = updateBook.Author
}
if updateBook.Author !=""{
	bookDetails.Publication = updateBook.Publication
}
db.Save(&bookDetails)
res,_:=json.Marshal(bookDetails)
w.Header().Set("Content-Type","pkglication/json")
w.WriteHeader(http.StatusOK)
w.Write(res)
}
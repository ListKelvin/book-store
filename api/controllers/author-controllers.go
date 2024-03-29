package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ListKelvin/book-store/api/models"
	"github.com/ListKelvin/book-store/api/utils"
	"github.com/gorilla/mux"
)


var NewBook models.Book 

func GetAuthors(w http.ResponseWriter, r *http.Request){

}

func GetAuthorById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId, err := vars["bookId"]
	// ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != false {
		fmt.Println("Error while converting")

	}

	bookDetails, _ := NewBook.GetBookById(bookId, db)


	response := models.Response{
    Data: bookDetails,  // Populate data based on request processing
    Message: "Success", // Or set an appropriate message
    Status:  "OK",
  }
	res, _ :=json.Marshal(response)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}
func  CreateAuthor(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b,_:= CreateBook.CreateBook(db)
	 response := models.Response{
    Data: b,  // Populate data based on request processing
    Message: "Success", // Or set an appropriate message
    Status:  "OK",
  }

	res, _ :=json.Marshal(response)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteAuthor(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	// ID, err := strconv.ParseInt(bookId, 0 ,0)


	book, er:= NewBook.DeleteBook(bookId , db)

	if er != nil {
		fmt.Println("Error in db: ", er.Error())
		
	}
	response := models.Response{
    Data: book,  // Populate data based on request processing
    Message: "Success", // Or set an appropriate message
    Status:  "OK",
  	}
	res, _ :=json.Marshal(response)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)

}

// need to fix
func UpdateBook(w http.ResponseWriter, r *http.Request){
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId, err := vars["bookId"]
	// ID, err := strconv.ParseInt(bookId, 0 ,0)
	if err != false {
		fmt.Println("Error while converting")

	}

	book , _ := NewBook.GetBookById(bookId,db)

	// if updateBook.Name != "" {
	// 	book.Name = updateBook.Name
	// }
	// if updateBook.Author != "" {
	// 	book.Name = updateBook.Author
	// }

	// if updateBook.Publication != "" {
	// 	book.Name = updateBook.Publication
	// }



	db.Save(&book)
	res, _ :=json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(res)
}
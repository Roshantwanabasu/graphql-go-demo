package repository

import (
	"log"

	"github.com/Roshantwanabasu/graphql-go-demo/graph/model"
	db "github.com/Roshantwanabasu/graphql-go-demo/internal/pkg/db/migrations/mysql"
)

func CreateBook(book model.Book) (int64, error) {
	stmt, err := db.Db.Prepare("Insert into Books(Title,AuthorID) values(?,?)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(book.Title, book.Author.ID)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return id, nil

}

func GetBookByID(id *string) (*model.Book, error) {
	stmt, err := db.Db.Prepare("select Books.ID,Books.Title,Authors.ID,Authors.FirstName,Authors.LastName from Books inner join Authors where Books.AuthorID = Authors.ID and Books.ID = ? ;")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	result, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var bookID, title, authorID, firstName, lastName string
	if result.Next() {
		err := result.Scan(&bookID, &title, &authorID, &firstName, &lastName)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	defer result.Close()
	book := &model.Book{
		ID:    bookID,
		Title: title,
		Author: &model.Author{
			ID:        authorID,
			FirstName: firstName,
			LastName:  lastName,
		},
	}
	return book, nil

}

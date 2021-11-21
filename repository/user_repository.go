package repository

import (
	"log"

	"github.com/Roshantwanabasu/graphql-go-demo/graph/model"
	db "github.com/Roshantwanabasu/graphql-go-demo/internal/pkg/db/migrations/mysql"
)

func CreateAuthor(author model.Author) (int64, error) {
	stmt, err := db.Db.Prepare("Insert into Authors(FirstName,LastName) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(author.FirstName, author.LastName)
	if err != nil {
		log.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row inserted")

	return id, nil

}

func GetAuthorById(id *string) (*model.Author, error) {
	stmt, err := db.Db.Prepare("Select * from Authors where id=?")
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
	defer result.Close()
	var author model.Author

	for result.Next() {
		err = result.Scan(&author.ID, &author.FirstName, &author.LastName)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	if err = result.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &author, nil
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Roshantwanabasu/graphql-go-demo/graph/generated"
	"github.com/Roshantwanabasu/graphql-go-demo/graph/model"
	"github.com/Roshantwanabasu/graphql-go-demo/repository"
)

func (r *mutationResolver) CreateBook(ctx context.Context, title string, author string) (*model.Book, error) {
	var book model.Book
	book.Title = title
	book.Author = &model.Author{
		ID: author,
	}
	id, err := repository.CreateBook(book)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	idStr := strconv.Itoa(int(id))
	createdBook, _ := repository.GetBookByID(&idStr)
	return createdBook, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, firstName string, lastName string) (*model.Author, error) {
	var author model.Author
	author.FirstName = firstName
	author.LastName = lastName
	id, _ := repository.CreateAuthor(author)
	return &model.Author{ID: strconv.FormatInt(id, 10), FirstName: author.FirstName, LastName: author.LastName}, nil
}

func (r *queryResolver) BookByID(ctx context.Context, id *string) (*model.Book, error) {
	book, err := repository.GetBookByID(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return book, nil
}

func (r *queryResolver) AllBooks(ctx context.Context) ([]*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AuthorByID(ctx context.Context, id *string) (*model.Author, error) {
	author, err := repository.GetAuthorById(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return author, nil
}

func (r *queryResolver) AllAuthors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

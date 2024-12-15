package repository

import (
	"TODOLIST/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
	GetBook(id string) (models.Book, error)
	GetAllBooks() ([]models.Book, error)
	CreateBook(models.Book) (models.Book, error)
	// UpdateBook(models.Book) (models.Book, error)
	// DeleteBook(id int) error
}

type BookRepositoryDB struct {
	DB *sqlx.DB
}

func NewBookRepositoryDB(db *sqlx.DB) *BookRepositoryDB {
	return &BookRepositoryDB{DB: db}
}
func (r *BookRepositoryDB) GetBook(id string) (models.Book, error) {
	var book models.Book

	err := r.DB.Get(&book, "select id, title, author from books where id = $1", id)

	if err != nil {
		log.Println(err)
		return book, err
	}
	return book, err
}

func (r *BookRepositoryDB) CreateBook(book models.Book) (models.Book, error) {
	_, err := r.DB.Exec("insert into books (title,author) values($1, $2)", book.Title, book.Author)

	if err != nil {
		log.Printf("Error creating book: %v", err)
		return book, err
	}
	return book, nil
}

func (r *BookRepositoryDB) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	err := r.DB.Select(&books, "Select id, title, author from books")

	if err != nil {
		log.Printf("Error fetching books: %v", err)
		return nil, err
	}
	return books, nil
}

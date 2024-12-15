package handlers

import (
	"TODOLIST/internal/models"
	"TODOLIST/internal/repository"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Repo repository.BookRepository
}

var books = []models.Book{}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{Repo: repo}
}
func (h *BookHandler) GetBooks(c *gin.Context) {

	books, err := h.Repo.GetAllBooks()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, books)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(200, book)
			return
		}
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var newBook models.Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	createdBook, err := h.Repo.CreateBook(newBook)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not create book"})
	}

	c.JSON(201, createdBook)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updateBook models.Book
	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for index, book := range books {
		if book.ID == id {
			books[index] = updateBook
			c.JSON(200, updateBook)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Book not found"})

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			c.JSON(200, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Book not found"})
}

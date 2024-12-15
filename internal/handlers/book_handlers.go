package handlers

import(
	"github.com/gin-gonic/gin"
)

type Book struct{
	ID string
	Title string
	Author string
}

var books = []Book{
	{"1","Book1", "Author1"},
	{"2","Book2", "Author2"},
	{"3","Book3", "Author3"},
}

func GetBooks(c *gin.Context){
	c.JSON(200, books)
}

func GetBookByID(c *gin.Context){
	id := c.Param("id")
	for _,book := range books{
		if book.ID == id{
			c.JSON(200, book)
			return
		}
	}
}

func CreateBook(c *gin.Context){
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(201, newBook)
}

func UpdateBook(c *gin.Context){
	id := c.Param("id")
	var updateBook Book
	if err := c.ShouldBindJSON(&updateBook); err != nil{
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	for index, book := range books{
		if book.ID == id{
			books[index] = updateBook
			c.JSON(200, updateBook)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Book not found"})

}

func DeleteBook(c *gin.Context){
	id := c.Param("id")
	for index, book := range books{
		if book.ID == id{
			books = append(books[:index], books[index+1:]...)
			c.JSON(200, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(404, gin.H{"error": "Book not found"})
}
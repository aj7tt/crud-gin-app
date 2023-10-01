package main

import (
    "github.com/gin-gonic/gin"
    "crud-gin-app/handlers"
)

func main() {
    router := gin.Default()
    
    // Define CRUD routes
    router.POST("/books", handlers.CreateBook)
    router.GET("/books", handlers.GetBooks)
    router.GET("/books/:id", handlers.GetBookByID)
    router.PUT("/books/:id", handlers.UpdateBook)
    router.DELETE("/books/:id", handlers.DeleteBook)
    
    router.Run(":8080")
}


// curl http://localhost:8080/books

//curl -X POST -H "Content-Type: application/json" -d "{\"title\":\"Book 1\",\"author\":\"Author 1\"}" http://localhost:8080/books

//curl -X DELETE http://localhost:8080/books/4

// curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Book Title","author":"Updated Author Name"}' http://localhost:8080/books/4

// curl -X PATCH -H "Content-Type: application/json" -d '{"title":"Updated Book Title"}' http://localhost:8080/books/4

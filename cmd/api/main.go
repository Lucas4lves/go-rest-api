package main

import (
	"log"
	"net/http"

	"github.com/Lucas4lves/go-rest-api/controller"
	"github.com/Lucas4lves/go-rest-api/db"
	"github.com/Lucas4lves/go-rest-api/repositories"
	"github.com/Lucas4lves/go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	dbConn, err := db.ConnectDb()

	if err != nil {
		log.Fatal(err)
	}

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "A-OK",
		})
	})

	productRepo, err := repositories.NewProductRepository(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	productUc, err := usecase.NewProductUseCase(*productRepo)
	if err != nil {
		log.Fatal(err)
	}

	productC, err := controller.NewProductController(*productUc)
	if err != nil {
		log.Fatal(err)
	}

	g.GET("/products", productC.GetProducts)
	g.POST("/create/product", productC.CreateProduct)
	g.GET("/product/:id", productC.GetProductById)
	g.Run(":8080")
}

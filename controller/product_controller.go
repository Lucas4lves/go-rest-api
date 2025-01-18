package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lucas4lves/go-rest-api/model"
	"github.com/Lucas4lves/go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	//TODO: Use Case
	ProductUseCase usecase.ProductUseCase
}

func NewProductController(pu usecase.ProductUseCase) (pc *ProductController, err error) {
	return &ProductController{
		ProductUseCase: pu,
	}, nil
}

func (p *ProductController) GetProducts(ctx *gin.Context) {
	products, err := p.ProductUseCase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *ProductController) CreateProduct(ctx *gin.Context) {

	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedProduct, err := p.ProductUseCase.CreateProduct(product)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (pc *ProductController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, &model.ProductResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "ID value must not be null",
		})

		return
	}

	productId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := pc.ProductUseCase.GetProductById(productId)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	if product != nil {
		ctx.JSON(http.StatusOK, product)
		return
	}

	ctx.JSON(http.StatusBadRequest, &model.ProductResponse{
		StatusCode: http.StatusBadRequest,
		Message:    "Such product does not exist in the database!",
	})
}

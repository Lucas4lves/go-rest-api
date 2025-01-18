package usecase

import (
	"fmt"

	"github.com/Lucas4lves/go-rest-api/model"
	"github.com/Lucas4lves/go-rest-api/repositories"
)

type ProductUseCase struct {
	//TODO: Repository
	Repository repositories.ProductRepository
}

func NewProductUseCase(rep repositories.ProductRepository) (u *ProductUseCase, err error) {
	return &ProductUseCase{
		Repository: rep,
	}, nil
}

func (pu *ProductUseCase) GetProducts() (products []model.Product, err error) {
	res, err := pu.Repository.GetProducts()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return res, nil
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.Repository.CreateProduct(product)

	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (*model.Product, error) {
	product, err := pu.Repository.GetProductById(id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return product, nil
}

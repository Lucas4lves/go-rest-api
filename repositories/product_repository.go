package repositories

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Lucas4lves/go-rest-api/model"
)

type ProductRepository struct {
	Conn *sql.DB
}

func NewProductRepository(conn *sql.DB) (*ProductRepository, error) {
	return &ProductRepository{
		Conn: conn,
	}, nil
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.Conn.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var product_list []model.Product
	var product_obj model.Product

	for rows.Next() {
		err = rows.Scan(
			&product_obj.ID,
			&product_obj.Name,
			&product_obj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		product_list = append(product_list, product_obj)
	}

	rows.Close()

	return product_list, nil

}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.Conn.Prepare("INSERT INTO product" +
		" (product_name, price)" +
		" VALUES ($1, $2) RETURNING id")

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return id, nil
}

func (pr *ProductRepository) GetProductById(id int) (*model.Product, error) {
	query, err := pr.Conn.Prepare("SELECT * FROM product " +
		"WHERE id=$1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product model.Product

	err = query.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}

	query.Close()

	return &product, nil
}

func (pr *ProductRepository) UpdateProduct(id int, partial model.Product) (*model.Product, error) {
	return nil, nil
}

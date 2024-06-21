package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Tountoun/ecom-api/types"
)

type Store struct {
	db *sql.DB
}


func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}


func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)

	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)", p.Name, p.Description, p.Image, p.Price, p.Quantity)
	return err
}

func (s *Store) GetProductsByIDs(ids []int) ([]types.Product, error) {
	// format request string
	argsIndicator := strings.Repeat(", ?", len(ids) - 1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", argsIndicator)

	args := make([]interface{}, len(ids))
	for i, v := range ids {
		args[i] = v
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	var products []types.Product

	for rows.Next() {
		product, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (s *Store) UpdateProduct(p types.Product) error {
	query := "UPDATE products SET name=?, description=?, image=?, price=?, quantity=? WHERE id=?"
	_, err := s.db.Exec(query, p.Name, p.Description, p.Image, p.Price, p.Quantity, p.ID)
	return err
}

func (s *Store) GetProductByID(id int) (types.Product, error) {
	product := new(types.Product)
	query := "SELECT * FROM products WHERE id=?"
	row := s.db.QueryRow(query, id)

	err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt)
	return *product, err
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}
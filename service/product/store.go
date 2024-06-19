package product

import (
	"database/sql"

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

func (s *Store) CreateProduct(p types.ProductPayload) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)", p.Name, p.Description, p.Image, p.Price, p.Quantity)
	return err
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
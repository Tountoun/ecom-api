package product

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Tountoun/ecom-api/types"
	"github.com/gorilla/mux"
)

func TestProductServiceHandlers(t *testing.T) {

	mockStore := &mockProductStore{}

	handler := NewHandler(mockStore)

	t.Run("get products should return an empty list", func(t *testing.T) {
		
		req, err := http.NewRequest(http.MethodGet, "/products", http.NoBody)
		if err != nil {
			t.Fatal(err)
		}
		
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleGetProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}

		if rr.Body == nil {
			t.Errorf("expected empty data but got nil")
		}

		resp := new([]types.ProductPayload)
		if err := json.NewDecoder(rr.Body).Decode(resp); err != nil {
			t.Errorf("error while decoding response body")
		}

		if len(*resp) != 0 {
			t.Errorf("expected empty data but got a data with %v elements", len(*resp))
		}
	})

	t.Run("should pass when creating a new product", func(t *testing.T) {
		payload := getPayload("t-shirt", "new shirt", "image.png", 453, 34)
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/products", handler.handleCreateProduct)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}


func getPayload(name, description, image string, price float64, quantity int) types.ProductPayload {
	return types.ProductPayload{
		Name: name,
		Description: description,
		Image: image,
		Price: price,
		Quantity: quantity,
	}
}

type mockProductStore struct {}

func (mock *mockProductStore) GetProducts() ([]types.Product, error) {
	return make([]types.Product, 0), nil
}

func (mock *mockProductStore) CreateProduct(payload types.Product) error {
	return nil
}

func (mock *mockProductStore) GetProductsByIDs(ids []int) ([]types.Product, error) {
	return nil, nil
}


func (mock *mockProductStore) UpdateProduct(product types.Product) error {
	return nil
}

func (mock *mockProductStore) GetProductByID(id int) (types.Product, error) {
	return types.Product{}, nil
}
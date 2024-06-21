package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Tountoun/ecom-api/service/cart"
	"github.com/Tountoun/ecom-api/service/order"
	"github.com/Tountoun/ecom-api/service/product"
	"github.com/Tountoun/ecom-api/service/user"
	"github.com/gorilla/mux"
)

// Holds server address and database
type APIServer struct {
	addr string
	db *sql.DB
}

// Return a new api server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr,
		db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	store := user.NewStore(s.db)
	userHandler := user.NewHandler(store)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)
	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, store)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Server starting listening on port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
package product

import (
	"net/http"

	"github.com/Tountoun/ecom-api/types"
	"github.com/Tountoun/ecom-api/utils"
	"github.com/gorilla/mux"
)


type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/products", h.handleGetProduct).Methods(http.MethodGet)
}


func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
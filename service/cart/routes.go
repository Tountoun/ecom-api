package cart

import (
	"fmt"
	"net/http"

	"github.com/Tountoun/ecom-api/service/auth"
	"github.com/Tountoun/ecom-api/types"
	"github.com/Tountoun/ecom-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.OrderStore
	productStore types.ProductStore
	userStore types.UserStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler {
		store: store,
		productStore: productStore,
		userStore: userStore,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", auth.WithJWTAuth(h.handleCheckout, h.userStore)).Methods(http.MethodPost)
}


func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutPayload
	userID := auth.GetUserIDFromContext(r.Context())

	if userID == -1 {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("unknown user"))
	}

	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// get products ids from the cart items
	productIds, err := getCartItemsProductIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// get products from db using their ids
	products, err := h.productStore.GetProductsByIDs(productIds)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	orderID, totalPrice, err := h.createOrder(products, cart.Items, userID)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any {
		"order_id": orderID,
		"total_price": totalPrice,
	})
}
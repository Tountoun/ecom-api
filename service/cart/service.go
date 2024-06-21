package cart

import (
	"fmt"

	"github.com/Tountoun/ecom-api/types"
)

func getCartItemsProductIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))

	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}
		productIds[i] = item.ProductID
	}

	return productIds, nil
}


func (h *Handler) createOrder(products []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	// check products availability in stock
	if err := checkIfCartItemInStock(items, productMap); err != nil {
		return 0, 0, err
	}

	// calculate the total price
	total := calculateTotalPrice(items, productMap)

	// reduce quantity of products in the database
	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity
		err := h.productStore.UpdateProduct(product)
		if err != nil {
			return 0, 0, err
		}
	}

	// create order
	order := types.Order{
		UserID: userID,
		Total: total,
		Status: "pending",
		Address: "Lomé",
	}

	orderID, err := h.store.CreateOrder(order)

	if err != nil {
		return 0, 0, err
	}

	// create order items
	for _, item := range items {
		product := productMap[item.ProductID]
		orderItem := types.OrderItem{
			OrderID: orderID,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			Price: product.Price,
		}
		if err := h.store.CreateOrderItem(orderItem); err != nil {
			return orderID, total, err
		}
	}
	return orderID, total, nil
}


func checkIfCartItemInStock(items []types.CartItem, productMap map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}
	for _, item := range items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product with id %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < item.Quantity {
			return fmt.Errorf("product %s is not available in the quantity requested", product.Name)
		}
	}
	
	return nil
}

func calculateTotalPrice(items []types.CartItem, productMap map[int]types.Product) float64 {
	var total float64

	for _, item := range items {
		product := productMap[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}
	
	return total
}
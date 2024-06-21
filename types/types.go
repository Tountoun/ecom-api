package types

import "time"

type RegisterPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginPayload struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type ProductPayload struct {
	Name 		string 		`json:"name" validate:"required"`
	Description string 		`json:"description" validate:"required"`
	Image 		string 		`json:"image" validate:"required"`
	Price 		float64 	`json:"price" validate:"required,min=0"`
	Quantity 	int 		`json:"quantity" validate:"required,min=1"`
}

type User struct {
	ID        int		`json:"id"`
	FirstName string 	`json:"firstName"`
	LastName  string 	`json:"lastName"`
	Email     string 	`json:"email"`
	Password  string 	`json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Description string 		`json:"description"`
	Image 		string 		`json:"image"`
	Price 		float64 	`json:"price"`
	Quantity 	int 		`json:"quantity"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}

type CartItem struct {
	ProductID 	int			`json:"productID"`
	Quantity 	int			`json:"quantity"`
}

type Order struct {
	ID 			int			`json:"id"`
	UserID 		int			`json:"userID"`
	Total 		float64		`json:"total"`
	Status 		string		`json:"status"`
	Address 	string		`json:"address"`
	CreatedAt 	time.Time	`json:"createdAt"`
}

type OrderItem struct {
	ID 			int			`json:"id"`
	OrderID		int			`json:"orderID"`
	ProductID	int			`json:"productID"`
	Quantity	int			`json:"quantity"`
	Price		float64		`json:"price"`
	CreatedAt 	time.Time	`json:"createdAt"`
}

type CartCheckoutPayload struct {
	Items []CartItem		`json:"items" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	CreateProduct(Product) error
	GetProductsByIDs(ids []int) ([]Product, error)
	UpdateProduct(Product) error
	GetProductByID(id int) (Product, error)
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}
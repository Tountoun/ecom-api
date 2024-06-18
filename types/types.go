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

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
}
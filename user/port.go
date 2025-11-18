package user

import (
	"ecommerece/domain"
	userHandler "ecommerece/rest/handlers/user"
)

type Service interface {
	userHandler.Service //embedings
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}

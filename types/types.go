package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	Id        int32     `json:"id"`
	FirstName string    `json:"first name"`
	LastName  string    `json:"Last Name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created at"`
}
type RegisterUserPayload struct {
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

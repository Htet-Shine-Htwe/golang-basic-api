package types

type RegisterUserPayload struct {
	FirstName string `json:"first name"`
	LastName  string `json:"last name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

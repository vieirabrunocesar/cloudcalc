package domain

const (
	// ADMINISTRADOR - Admin type user acess.
	ADMINISTRADOR = 1
	// CONSULTOR - Consultor type user acess.
	CONSULTOR = 2
)

// User login application.
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType int    `json:"userType"`
}

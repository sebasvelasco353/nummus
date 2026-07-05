package users

type User struct {
	ID       string
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Name     string `binding:"required"`
	LastName string
}

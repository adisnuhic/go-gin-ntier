package requests

type Registration struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

package requests

type RefreshToken struct {
	Token string `json:"token" binding:"required"`
}

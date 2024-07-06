package requests

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"john.doe@example.com"`
	Password string `json:"password" binding:"required" example:"123456"`
}

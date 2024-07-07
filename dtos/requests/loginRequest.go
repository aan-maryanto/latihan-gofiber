package requests

type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"john.doe@example.com" validate:"required,email,max=100"`
	Password string `json:"password" binding:"required" example:"123456" validate:"required,min=8"`
}

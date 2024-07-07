package requests

type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe" validate:"required, max=100"`
	Email    string `json:"email" binding:"required" example:"john.doe@example.com" validate:"required,email,max=100"`
	Password string `json:"password" binding:"required" example:"123456", validate:"required,min=8"`
}

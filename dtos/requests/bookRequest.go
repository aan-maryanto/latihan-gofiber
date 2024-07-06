package requests

type BookRequest struct {
	Title       string `json:"title" binding:"required" example:"Jane Doe"`
	Author      string `json:"author" binding:"required" example:"John Doe"`
	Publisher   string `json:"publisher" binding:"required" example:"https://www.google.com"`
	Category    string `json:"category" binding:"required" example:"Category"`
	Year        int    `json:"year" binding:"required" example:"2020"`
	ISBN        string `json:"isbn" binding:"required" example:"978-978"`
	Description string `json:"description" binding:"required" example:"Jane Doe"`
	Image       string `json:"image" binding:"required" example:"https://www.google.com"`
}

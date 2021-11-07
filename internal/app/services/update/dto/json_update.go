package dto

type UpdateRequest struct {
	ID          string `json:"id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required,max=100"`
	Description string `json:"description"`
	Date        string `json:"date" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Image       string `json:"image" validate:"base64"`
}

type UpdateResponse struct {
	ErrorCode string `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

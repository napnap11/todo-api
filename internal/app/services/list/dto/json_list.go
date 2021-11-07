package dto

import "github.com/napnap11/todo-api/internal/pkg/models"

type ListRequest struct {
	SortBy      string `json:"sort_by"`
	SortType    string `json:"sort_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ListResponse struct {
	Todos []models.Todo
	ErrorCode string `json:"error_code"`
	ErrorDesc string `json:"error_desc"`
}

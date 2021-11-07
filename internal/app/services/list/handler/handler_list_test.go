package handler

import (
	"bytes"
	"encoding/json"
	"github.com/napnap11/todo-api/internal/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/napnap11/todo-api/internal/app/services/list/dto"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
}

func (mockService) List(req dto.ListRequest) dto.ListResponse {
	return dto.ListResponse{
		Todos:     []models.Todo{},
		ErrorCode: "00",
		ErrorDesc: "success",
	}
}

func TestHandlerListSuccessCase(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	jsonReq := dto.ListRequest{
		SortBy:      "title",
		SortType:    "desc",
		Title:       "",
		Description: "",
	}
	rawReq, err := json.Marshal(jsonReq)
	if err != nil {
		t.Error(err)
	}
	req, _ := http.NewRequest("POST", "/v1/list", bytes.NewBuffer(rawReq))
	s := mockService{}
	h := NewHandler(s)
	r.Handle("POST", "/v1/list", h.List)
	r.ServeHTTP(w, req)

	var resp dto.ListResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}

	expected := dto.ListResponse{
		Todos:     []models.Todo{},
		ErrorCode: "00",
		ErrorDesc: "success",
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.Todos, resp.Todos)
	assert.Equal(t, expected.ErrorCode, resp.ErrorCode)
	assert.Equal(t, expected.ErrorDesc, resp.ErrorDesc)
}

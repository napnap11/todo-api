package handler

import (
	"bytes"
	"encoding/json"
	"github.com/napnap11/todo-api/internal/app/services/create/dto"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
}

func (mockService) Create(req dto.CreateRequest) dto.CreateResponse {
	return dto.CreateResponse{
		ErrorCode: "00",
		ErrorDesc: "success",
	}
}

func TestHandlerCreateSuccessCase(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	jsonReq := dto.CreateRequest{
		ID:     "123e4567-e89b-12d3-a456-426614174000",
		Title:  "test",
		Date:   time.Now().Format(time.RFC3339),
		Status: "IN_PROGRESS",
		Image:  "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",
	}
	rawReq, err := json.Marshal(jsonReq)
	if err != nil {
		t.Error(err)
	}
	req, _ := http.NewRequest("POST", "/v1/create", bytes.NewBuffer(rawReq))
	s := mockService{}
	h := NewHandler(s)
	r.Handle("POST", "/v1/create", h.Create)
	r.ServeHTTP(w, req)

	var resp dto.CreateResponse
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	if err != nil {
		t.Error(err)
	}

	expected := dto.CreateResponse{
		ErrorCode: "00",
		ErrorDesc: "success",
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected.ErrorCode, resp.ErrorCode)
	assert.Equal(t, expected.ErrorDesc, resp.ErrorDesc)
}

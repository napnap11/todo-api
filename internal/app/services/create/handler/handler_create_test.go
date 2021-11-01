package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
}

func TestHandlerCreateSuccessCase(t *testing.T) {
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)
	req, _ := http.NewRequest("POST", "/v1/create", nil)
	s := mockService{}
	h := NewHandler(s)
	r.Handle("POST", "/v1/create", h.Create)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "", w.Body.String())
}

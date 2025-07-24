package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloHandler(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Act
	HelloHandler(c)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	expected := `{"message":"hello from v2 (new logic)"}`
	assert.JSONEq(t, expected, w.Body.String())
}

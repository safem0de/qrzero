// internal\04_api\v1\handler\custable_test.go

package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"qrzero/internal/01_entity"
)

// MockCustableService implements application.CustableService
type MockCustableService struct {
	mock.Mock
}

func (m *MockCustableService) GetRecentActiveCustomers(ctx context.Context) ([]entity.Custable, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Custable), args.Error(1)
}

func TestGetRecentActiveCustomers_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(MockCustableService)
	expected := []entity.Custable{
		{BillerID: "001", AccountNum: "123456", CompanyBank: "SCB", Name: "Alice", CustomerStatus: 0},
	}
	mockSvc.On("GetRecentActiveCustomers", mock.Anything).Return(expected, nil)

	handler := NewCustableHandler(mockSvc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/custable", nil)

	handler.GetRecentActiveCustomers(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Alice")
	mockSvc.AssertExpectations(t)
}

func TestGetRecentActiveCustomers_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(MockCustableService)

	mockSvc.On("GetRecentActiveCustomers", mock.Anything).
		Return(([]entity.Custable)(nil), errors.New("DB connection failed"))

	handler := NewCustableHandler(mockSvc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/v1/custable", nil)

	handler.GetRecentActiveCustomers(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "DB connection failed")
	mockSvc.AssertExpectations(t)
}

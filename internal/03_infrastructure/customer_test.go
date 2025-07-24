// internal/03_infrastructure/customer_test.go

package infrastructure

import (
	"errors"
	"os"
	"testing"
	"qrzero/internal/01_entity"
	"qrzero/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
)

func TestCheckFileExist_FileExists(t *testing.T) {
	// Arrange: สร้างไฟล์ temp
	testFile := "test_file.tmp"
	path := "./qrcode/" + testFile
	os.MkdirAll("./qrcode", 0755)
	os.WriteFile(path, []byte("hello"), 0644)
	defer os.Remove(path)

	s := &customerServiceImpl{}

	// Act
	exists := s.CheckFileExist(testFile)

	// Assert
	assert.True(t, exists)
}

func TestCheckFileExist_FileNotExists(t *testing.T) {
	testFile := "not_exist_file.tmp"
	s := &customerServiceImpl{}

	exists := s.CheckFileExist(testFile)
	assert.False(t, exists)
}

func TestPublishQRJob_Success(t *testing.T) {
	mockRabbit := new(mocks.MockRabbitMQClient)
	service := NewCustomerService(mockRabbit)

	req := entity.QRJobRequest{
		BillerID:    "BILLER1",
		AccountNum:  "ACC123",
		CompanyBank: "KBank",
		Amount:      "100",
		FilePath:    "test.png",
	}

	// set expectation
	mockRabbit.On("Publish", "qr_job", mock.AnythingOfType("[]uint8")).Return(nil)

	err := service.PublishQRJob(req)
	assert.NoError(t, err)
	mockRabbit.AssertExpectations(t)
}

func TestPublishQRJob_Fail(t *testing.T) {
	mockRabbit := new(mocks.MockRabbitMQClient)
	service := NewCustomerService(mockRabbit)
	req := entity.QRJobRequest{}

	mockRabbit.On("Publish", "qr_job", mock.AnythingOfType("[]uint8")).Return(errors.New("publish error"))

	err := service.PublishQRJob(req)
	assert.EqualError(t, err, "publish error")
	mockRabbit.AssertExpectations(t)
}

// test\mocks\mock_rabbitmq.go

package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockRabbitMQClient struct {
	mock.Mock
}

func (m *MockRabbitMQClient) Publish(queue string, body []byte) error {
	args := m.Called(queue, body)
	return args.Error(0)
}
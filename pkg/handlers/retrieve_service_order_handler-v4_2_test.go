package handler_v4_2

import (
	"context"
	"errors"
	"github.com/MxelA/tmf-service-go/pkg/mocks"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRetrieveServiceOrderHandler_Success(t *testing.T) {
	// 1. Setup mock
	mockRepo := &mocks.MockServiceOrderRepository{
		GetByIDFunc: func(ctx context.Context, id string, fields *string) (*models.ServiceOrder, error) {
			assert.Equal(t, "test-id", id) // Verify correct ID
			return &models.ServiceOrder{ID: "test-id"}, nil
		},
	}

	// 2. Create handler with mock
	handler := NewServiceOrderHandler(mockRepo)

	// 3. Create request parameters
	fieldsParam := "id,status"
	req := service_order.RetrieveServiceOrderParams{
		ID:          "test-id",
		Fields:      &fieldsParam,
		HTTPRequest: &http.Request{},
	}

	// 4. Execute handler
	response := handler.RetrieveServiceOrderHandler(req)

	// 5. Assert response type and content
	okResponse, ok := response.(*service_order.RetrieveServiceOrderOK)
	assert.True(t, ok, "Expected OK response")
	assert.Equal(t, "test-id", okResponse.Payload.ID)
}

func TestRetrieveServiceOrderHandler_Error(t *testing.T) {
	// 1. Setup mock to return error
	mockRepo := &mocks.MockServiceOrderRepository{
		GetByIDFunc: func(ctx context.Context, id string, fields *string) (*models.ServiceOrder, error) {
			return nil, errors.New("database connection failed")
		},
	}

	// 2. Create handler
	handler := NewServiceOrderHandler(mockRepo)

	// 3. Create request
	req := service_order.RetrieveServiceOrderParams{
		ID:          "test-id",
		HTTPRequest: &http.Request{},
	}

	// 4. Execute handler
	response := handler.RetrieveServiceOrderHandler(req)

	// 5. Verify error response
	errorResponse, ok := response.(*service_order.RetrieveServiceOrderInternalServerError)
	assert.True(t, ok, "Expected 500 error")
	assert.Equal(t, "Internal server error", errorResponse.Payload.Message)
	assert.Equal(t, "500", *errorResponse.Payload.Code)
}

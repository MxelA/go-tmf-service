package handler_v4_2

import (
	"context"
	"github.com/MxelA/tmf-service-go/pkg/mocks"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"net/url"
	"testing"
)

func TestListServiceOrderHandler_Success(t *testing.T) {
	// 1. Setup mock
	mockRepo := &mocks.MockServiceOrderRepository{
		GetAllPaginateFunc: func(ctx context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error) {
			var total = int64(1)
			return []*models.ServiceOrder{{ID: "1"}}, &total, nil
		},
	}

	// 2. Create handler with mock
	handler := NewServiceOrderHandler(mockRepo)
	offset := int64(0)
	limit := int64(10)
	fields := "id,state"
	// 3. Create request parameters
	req := service_order.ListServiceOrderParams{
		Fields: &fields,
		Offset: &offset,
		Limit:  &limit,
		HTTPRequest: &http.Request{
			URL: &url.URL{
				RawQuery: "state=pending",
			},
		},
	}

	// 4. Execute handler
	response := handler.ListServiceOrderHandler(req)
	assert.Equal(t, int64(1), int64(1))
	// 5. Assert response type and content
	assert.IsType(t, &service_order.ListServiceOrderOK{}, response)

	okResponse := response.(*service_order.ListServiceOrderOK)
	assert.Equal(t, int64(1), okResponse.XTotalCount)
	assert.Equal(t, int64(1), okResponse.XResultCount)
	assert.Len(t, okResponse.Payload, 1)

	// Assert repository calls
	expectedFilter := utils.BuildMongoFilter(req.HTTPRequest.URL.Query())
	filter := bson.M{}
	filter["state"] = bson.M{"$eq": "pending"}
	assert.Equal(t, expectedFilter, filter)
}

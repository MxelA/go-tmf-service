package handlers_test

import (
	"context"
	database "github.com/MxelA/tmf-service-go/pkg/config"
	handler_v4_2 "github.com/MxelA/tmf-service-go/pkg/handlers"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"testing"
)

type MockMongoDB struct {
	mock.Mock
}

func (m *MockMongoDB) GetMongoInstance() database.MongoInstance {
	args := m.Called()
	return args.Get(0).(database.MongoInstance)
}

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func TestCreateServiceOrderHandler_Success(t *testing.T) {

	mockCollection := new(MockCollection)

	// Mock InsertOne to return a successful result
	insertResult := &mongo.InsertOneResult{InsertedID: "123"}
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(insertResult, nil)

	// Mock FindOne to return a mock document
	mockSingleResult := &mongo.SingleResult{}
	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

	mockMongoDB := new(MockMongoDB)
	expectedInstance := database.MongoInstance{
		Client: nil,
		Db:     nil,
	}
	mockMongoDB.On("GetMongoInstance").Return(expectedInstance)

	// Create a mock request
	req := service_order.CreateServiceOrderParams{
		HTTPRequest:  &http.Request{},
		ServiceOrder: &models.ServiceOrderCreate{
			// Populate with test data
		},
	}

	// Use the mock in your test
	result := mockMongoDB.GetMongoInstance()
	assert.Equal(t, expectedInstance, result)

	// Assert that the method was called
	mockMongoDB.AssertCalled(t, "GetMongoInstance")

	// Call the handler
	response := handler_v4_2.CreateServiceOrderHandler(req)
	// Assert that the response is as expected
	assert.IsType(t, &service_order.CreateServiceOrderCreated{}, response)
}

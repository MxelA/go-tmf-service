package handler_v4_2_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

// MockMongoCollection is a mock implementation of the MongoDB collection.
type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

// MockSingleResult is a mock implementation of the MongoDB SingleResult.
type MockSingleResult struct {
	mock.Mock
}

func (m *MockSingleResult) Decode(v interface{}) error {
	args := m.Called(v)
	return args.Error(0)
}

func TestCreateServiceOrderHandler(t *testing.T) {
	// Mock MongoDB collection
	mockCollection := new(MockMongoCollection)
	mockSingleResult := new(MockSingleResult)

	// Mock the InsertOne method
	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "123"}, nil)

	// Mock the FindOne method
	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

	// Mock the Decode method
	mockSingleResult.On("Decode", mock.Anything).Return(nil)

	// Replace the actual MongoDB collection with the mock
	//database.GetMongoInstance = func() *database.MongoInstance {
	//	return &database.MongoInstance{
	//		Db: &mongo.Database{
	//			Collection: func(name string) *mongo.Collection {
	//				return mockCollection
	//			},
	//		},
	//	}
	//}
	//
	//// Create a mock request
	//req := service_order.CreateServiceOrderParams{
	//	HTTPRequest:  &http.Request{},
	//	ServiceOrder: &models.ServiceOrderCreate{},
	//}
	//
	//// Call the handler
	//response := handler_v4_2.CreateServiceOrderHandler(req)

	// Assertions
	//require.IsType(t, &service_order.CreateServiceOrderCreated{}, response)
	mockCollection.AssertExpectations(t)
	mockSingleResult.AssertExpectations(t)
}

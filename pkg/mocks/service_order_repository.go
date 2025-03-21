package mocks

import (
	"context"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockServiceOrderRepository struct {
	GetByIDFunc        func(ctx context.Context, id string, fields *string) (*models.ServiceOrder, error)
	GetAllPaginateFunc func(context context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error)
	CreateFunc         func(context context.Context, serviceOrder *models.ServiceOrderCreate) (*mongo.InsertOneResult, error)
	MergePatchFunc     func(context context.Context, id string, serviceOrder *models.ServiceOrderUpdate) (bool, error)
	DeleteFunc         func(context context.Context, id string) (bool, error)
}

func (m *MockServiceOrderRepository) GetByID(
	ctx context.Context,
	id string,
	fields *string,
) (*models.ServiceOrder, error) {
	return m.GetByIDFunc(ctx, id, fields)
}

func (m *MockServiceOrderRepository) GetAllPaginate(
	ctx context.Context,
	queryParams bson.M, selectFields *string, offset *int64, limit *int64,
) ([]*models.ServiceOrder, *int64, error) {
	return m.GetAllPaginateFunc(ctx, queryParams, selectFields, offset, limit)
}

func (m *MockServiceOrderRepository) Create(
	ctx context.Context,
	serviceOrder *models.ServiceOrderCreate,
) (*mongo.InsertOneResult, error) {
	return m.CreateFunc(ctx, serviceOrder)
}

func (m *MockServiceOrderRepository) MergePatch(
	ctx context.Context,
	id string, serviceOrder *models.ServiceOrderUpdate,
) (bool, error) {
	return m.MergePatchFunc(ctx, id, serviceOrder)
}

func (m *MockServiceOrderRepository) Delete(
	ctx context.Context,
	id string,
) (bool, error) {
	return m.DeleteFunc(ctx, id)
}

package repository

import (
	"context"
	"errors"
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type ServiceOrderRepository interface {
	GetByID(id string, selectFields *string) (*models.ServiceOrder, error)
	GetAllPaginate(queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, int64, error)
	GetAll() ([]*models.ServiceOrder, error)
	Create(serviceOrder *models.ServiceOrderCreate) (*mongo.InsertOneResult, error)
	Update(user *models.ServiceOrder) error
	Delete(id string) error
}

type MongoServiceOrderRepository struct {
	MongoInstance database.MongoInstance
	Context       context.Context
}

func (repo *MongoServiceOrderRepository) GetByID(id string, selectFields *string) (*models.ServiceOrder, error) {

	serviceOrderId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("Id is not valid")
	}

	collection := repo.MongoInstance.Db.Collection("serviceOrder")
	filter := bson.D{{Key: "_id", Value: serviceOrderId}}

	// Apply projection if set
	findOptions := options.FindOne()

	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	record := collection.FindOne(repo.Context, filter, findOptions)

	retrieveServiceOrder := models.ServiceOrder{}
	err = record.Decode(&retrieveServiceOrder)

	if err != nil {
		return nil, err
	}

	return &retrieveServiceOrder, nil
}

func (repo *MongoServiceOrderRepository) GetAllPaginate(queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error) {

	offset, limit = utils.ValidatePaginationParams(offset, limit)
	collection := repo.MongoInstance.Db.Collection("serviceOrder")

	findOptions := &options.FindOptions{ // Find options
		Skip:  offset,
		Limit: limit,
	}

	// Fields Projection
	fieldProjection := utils.GerFieldsProjection(selectFields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	// Get list of service orders

	records, err := collection.Find(repo.Context, queryParams, findOptions)
	if err != nil {
		return nil, nil, err
	}

	retrieveServiceOrders := []*models.ServiceOrder{}
	for records.Next(repo.Context) {
		var serviceOrder = models.ServiceOrder{}
		if err := records.Decode(&serviceOrder); err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		retrieveServiceOrders = append(retrieveServiceOrders, &serviceOrder) // Append pointer
	}

	totalCount, err := collection.CountDocuments(repo.Context, queryParams)

	if err != nil {
		return nil, nil, err
	}

	return retrieveServiceOrders, &totalCount, nil
}

func (repo *MongoServiceOrderRepository) Create(serviceOrder *models.ServiceOrderCreate) (*mongo.InsertOneResult, error) {
	mg := repo.MongoInstance
	collection := mg.Db.Collection("serviceOrder")

	insertResult, err := collection.InsertOne(repo.Context, serviceOrder)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (repo MongoServiceOrderRepository) Delete(id string) (bool, error) {
	serviceOrderId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	collection := repo.MongoInstance.Db.Collection("serviceOrder")
	filter := bson.D{{Key: "_id", Value: serviceOrderId}}
	deleteRecord, err := collection.DeleteOne(repo.Context, filter)

	if err != nil {
		return false, err
	}

	if deleteRecord.DeletedCount == 0 {
		return false, errors.New("Delete record with ID:" + id + " not success")
	}

	return true, nil
}

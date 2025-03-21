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
	GetByID(context context.Context, id string, selectFields *string) (*models.ServiceOrder, error)
	GetAllPaginate(context context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error)
	Create(context context.Context, serviceOrder *models.ServiceOrderCreate) (*mongo.InsertOneResult, error)
	MergePatch(context context.Context, id string, serviceOrder *models.ServiceOrderUpdate) (bool, error)
	Delete(context context.Context, id string) (bool, error)
}

type MongoServiceOrderRepository struct {
	MongoInstance database.MongoInstance
}

func (repo *MongoServiceOrderRepository) GetByID(context context.Context, id string, selectFields *string) (*models.ServiceOrder, error) {

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

	record := collection.FindOne(context, filter, findOptions)

	retrieveServiceOrder := models.ServiceOrder{}
	err = record.Decode(&retrieveServiceOrder)

	if err != nil {
		return nil, err
	}

	return &retrieveServiceOrder, nil
}

func (repo *MongoServiceOrderRepository) GetAllPaginate(context context.Context, queryParams bson.M, selectFields *string, offset *int64, limit *int64) ([]*models.ServiceOrder, *int64, error) {

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

	records, err := collection.Find(context, queryParams, findOptions)
	if err != nil {
		return nil, nil, err
	}

	retrieveServiceOrders := []*models.ServiceOrder{}
	for records.Next(context) {
		var serviceOrder = models.ServiceOrder{}
		if err := records.Decode(&serviceOrder); err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		retrieveServiceOrders = append(retrieveServiceOrders, &serviceOrder) // Append pointer
	}

	totalCount, err := collection.CountDocuments(context, queryParams)

	if err != nil {
		return nil, nil, err
	}

	return retrieveServiceOrders, &totalCount, nil
}

func (repo *MongoServiceOrderRepository) Create(context context.Context, serviceOrder *models.ServiceOrderCreate) (*mongo.InsertOneResult, error) {
	mg := repo.MongoInstance
	collection := mg.Db.Collection("serviceOrder")

	insertResult, err := collection.InsertOne(context, serviceOrder)

	if err != nil {
		return nil, err
	}

	return insertResult, nil
}

func (repo *MongoServiceOrderRepository) Delete(context context.Context, id string) (bool, error) {
	serviceOrderId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	collection := repo.MongoInstance.Db.Collection("serviceOrder")
	filter := bson.D{{Key: "_id", Value: serviceOrderId}}
	deleteRecord, err := collection.DeleteOne(context, filter)

	if err != nil {
		return false, err
	}

	if deleteRecord.DeletedCount == 0 {
		return false, errors.New("Delete record with ID:" + id + " not success")
	}

	return true, nil
}

func (repo *MongoServiceOrderRepository) MergePatch(context context.Context, id string, serviceOrder *models.ServiceOrderUpdate) (bool, error) {
	serviceOrderId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	collection := repo.MongoInstance.Db.Collection("serviceOrder")

	filter := bson.M{"_id": serviceOrderId}
	update := bson.M{"$set": serviceOrder}
	record := collection.FindOneAndUpdate(context, filter, update)

	updateServiceOrder := models.ServiceOrder{}
	err = record.Decode(&updateServiceOrder)
	if err != nil {
		return false, err
	}

	return true, nil
}

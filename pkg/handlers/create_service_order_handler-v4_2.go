package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/repository"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateServiceOrderHandler(req service_order.CreateServiceOrderParams) middleware.Responder {

	mongoServiceOrderRepo := repository.MongoServiceOrderRepository{
		MongoInstance: database.GetMongoInstance(),
		Context:       req.HTTPRequest.Context(),
	}

	insertResult, err := mongoServiceOrderRepo.Create(req.ServiceOrder)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewCreateServiceOrderInternalServerError().WithPayload(&errModel)
	}

	// Get mongo document
	id := insertResult.InsertedID.(primitive.ObjectID).Hex()
	retrieveServiceOrder, err := mongoServiceOrderRepo.GetByID(id, nil)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		errModel := models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewRetrieveServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewCreateServiceOrderCreated().WithPayload(retrieveServiceOrder)

}

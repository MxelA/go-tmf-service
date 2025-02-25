package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func RetrieveServiceOrderHandler(req service_order.RetrieveServiceOrderParams) middleware.Responder {

	serviceOrderId, err := primitive.ObjectIDFromHex(req.ID)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		return service_order.NewCreateServiceOrderInternalServerError().WithPayload(&errModel)
	}

	// Get mongoDB instance
	mg := database.GetMongoInstance()
	collection := mg.Db.Collection("serviceOrder")
	filter := bson.D{{Key: "_id", Value: serviceOrderId}}

	// Apply projection if set
	findOptions := options.FindOne()

	fieldProjection := utils.GerFieldsProjection(req.Fields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	record := collection.FindOne(req.HTTPRequest.Context(), filter, findOptions)

	// Decode into bson.M first
	response, err := utils.ConvertBsonMToMinimalJSONResponse(*record)

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

	// Return structured response
	return service_order.NewRetrieveServiceOrderOKRaw().WithPayload(response)

	// Map record from mongoDB to Response model
	//retrieveServiceOrder := models.ServiceOrder{}
	//err = record.Decode(&retrieveServiceOrder)
	//if err != nil {
	//	errCode := "500"
	//	reason := err.Error()
	//	var errModel = models.Error{
	//		Reason:  &reason,
	//		Code:    &errCode,
	//		Message: "Internal server error",
	//	}
	//	log.Println(err)
	//	return service_order.NewRetrieveServiceOrderInternalServerError().WithPayload(&errModel)
	//}
	//
	//return service_order.NewCreateServiceOrderCreated().WithPayload(&retrieveServiceOrder)
}

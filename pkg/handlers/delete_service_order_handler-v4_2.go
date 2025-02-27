package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func DeleteServiceOrderHandler(req service_order.DeleteServiceOrderParams) middleware.Responder {

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

	record, err := collection.DeleteOne(req.HTTPRequest.Context(), filter)

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

	if record.DeletedCount == 0 {
		return service_order.NewDeleteServiceOrderBadRequest()
	} else {
		return service_order.NewDeleteServiceOrderNoContent()
	}
}

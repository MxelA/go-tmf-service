package controllers_tmf_service_v4_2

import (
	"fmt"
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateServiceOrderControllerHandler(req service_order.CreateServiceOrderParams) middleware.Responder {

	var res models.ServiceOrder

	// Copy data from create service order model to response service model
	err := copier.Copy(&res, &req.ServiceOrder)
	if err != nil {

		errCode := "500"
		reason := "Error during Copy CreateServiceOrder type to ServiceOrder"

		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewCreateServiceOrderInternalServerError().WithPayload(&errModel)
	}

	mg := database.GetMongoInstance()
	collection := mg.Db.Collection("serviceOrder")

	//res.ID = "test123"
	insertResult, err := collection.InsertOne(req.HTTPRequest.Context(), res)
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

	if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
		res.ID = oid.Hex() // Convert ObjectID to its string representation
	} else {
		res.ID = fmt.Sprintf("%v", insertResult.InsertedID) // Handle other types of IDs if not ObjectID
	}

	return service_order.NewCreateServiceOrderCreated().WithPayload(&res)
}

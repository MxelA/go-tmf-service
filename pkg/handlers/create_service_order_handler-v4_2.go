package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func CreateServiceOrderHandler(req service_order.CreateServiceOrderParams) middleware.Responder {
	doc := &models.ServiceOrder{}

	// Copy data from create service order model to response service model
	err := copier.Copy(&doc, &req.ServiceOrder)
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

	// Insert data to DB
	mg := database.GetMongoInstance()
	collection := mg.Db.Collection("serviceOrder")

	insertResult, err := collection.InsertOne(req.HTTPRequest.Context(), doc)
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

	//if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
	//	doc.ID = oid.Hex() // Convert ObjectID to its string representation
	//} else {
	//	doc.ID = fmt.Sprintf("%v", insertResult.InsertedID) // Handle other types of IDs if not ObjectID
	//}

	filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}
	record := collection.FindOne(req.HTTPRequest.Context(), filter)

	//var raw bson.Raw
	//err = record.Decode(&raw)
	//if err != nil {
	//	errCode := "500"
	//	reason := err.Error()
	//	var errModel = models.Error{
	//		Reason:  &reason,
	//		Code:    &errCode,
	//		Message: "Internal server error",
	//	}
	//	return service_order.NewCreateServiceOrderInternalServerError().WithPayload(&errModel)
	//}
	//log.Println("Raw data from MongoDB:", raw)

	createdServiceOrder := models.ServiceOrder{}
	err = record.Decode(&createdServiceOrder)
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
	//utils.PrettyPrint(createdServiceOrder)

	return service_order.NewCreateServiceOrderCreated().WithPayload(&createdServiceOrder)
}

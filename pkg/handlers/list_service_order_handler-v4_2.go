package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ListServiceOrderHandler(req service_order.ListServiceOrderParams) middleware.Responder {

	req.Offset, req.Limit = utils.ValidatePaginationParams(req.Offset, req.Limit)

	// Get mongoDB instance
	mg := database.GetMongoInstance()
	collection := mg.Db.Collection("serviceOrder")

	findOptions := &options.FindOptions{ // Find options
		Limit: req.Limit,
		Skip:  req.Offset,
	}
	// Fields Projection
	fieldProjection := utils.GerFieldsProjection(req.Fields)
	if len(fieldProjection) > 0 { // Only set projection if fields are provided
		findOptions.SetProjection(fieldProjection)
	}

	// Get list of service orders
	filter := utils.BuildMongoFilter(req.HTTPRequest.URL.Query())
	records, err := collection.Find(req.HTTPRequest.Context(), filter, findOptions)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewListServiceOrderInternalServerError().WithPayload(&errModel)
	}

	retrieveServiceOrders := []*models.ServiceOrder{}
	for records.Next(req.HTTPRequest.Context()) {
		var serviceOrder = models.ServiceOrder{}
		if err := records.Decode(&serviceOrder); err != nil {
			log.Println("Error decoding document:", err)
			continue
		}
		retrieveServiceOrders = append(retrieveServiceOrders, &serviceOrder) // Append pointer
	}

	// Get total count of documents
	countServiceOrders, err := collection.CountDocuments(req.HTTPRequest.Context(), filter)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewListServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewListServiceOrderOK().
		WithXTotalCount(countServiceOrders).
		WithXResultCount(int64(len(retrieveServiceOrders))).
		WithPayload(retrieveServiceOrders)
}

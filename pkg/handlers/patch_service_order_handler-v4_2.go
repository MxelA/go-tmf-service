package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/repository"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"strings"
)

func PatchServiceOrderHandler(req service_order.PatchServiceOrderParams) middleware.Responder {
	contentType := req.HTTPRequest.Header["Content-Type"][0]

	if strings.Contains(contentType, "application/json-patch+json") {
		return processJsonPatch(req)
	} else if strings.Contains(contentType, "application/merge-patch+json") {
		return processMergePatch(req)
	}

	errCode := "422"
	reason := "Unsupported Media Type"
	errModel := models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported media type" + contentType,
	}

	return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
}

func processJsonPatch(req service_order.PatchServiceOrderParams) middleware.Responder {
	errCode := "422"
	reason := "Unsupported media type "
	var errModel = models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported 'application/json-patch+json' media type ",
	}
	return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
}

func processMergePatch(req service_order.PatchServiceOrderParams) middleware.Responder {
	//serviceOrderId, err := primitive.ObjectIDFromHex(req.ID)
	//
	//if err != nil {
	//	errCode := "500"
	//	reason := err.Error()
	//	var errModel = models.Error{
	//		Reason:  &reason,
	//		Code:    &errCode,
	//		Message: "Internal server error",
	//	}
	//	return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	//}
	//
	//// Get mongoDB instance
	//mg := database.GetMongoInstance()
	//collection := mg.Db.Collection("serviceOrder")
	//
	//filter := bson.M{"_id": serviceOrderId}
	//update := bson.M{"$set": req.ServiceOrder}
	//record := collection.FindOneAndUpdate(req.HTTPRequest.Context(), filter, update)
	//
	//updateServiceOrder := models.ServiceOrder{}
	//err = record.Decode(&updateServiceOrder)
	mongoServiceOrderRepo := repository.MongoServiceOrderRepository{
		MongoInstance: database.GetMongoInstance(),
		Context:       req.HTTPRequest.Context(),
	}

	_, err := mongoServiceOrderRepo.MergePatch(req.ID, req.ServiceOrder)
	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
	}

	retrieveServiceOrder, err := mongoServiceOrderRepo.GetByID(req.ID, nil)

	return service_order.NewPatchServiceOrderOK().WithPayload(retrieveServiceOrder)
}

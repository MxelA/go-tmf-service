package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/repository"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func DeleteServiceOrderHandler(req service_order.DeleteServiceOrderParams) middleware.Responder {
	mongoServiceOrderRepo := repository.MongoServiceOrderRepository{
		MongoInstance: database.GetMongoInstance(),
		Context:       req.HTTPRequest.Context(),
	}

	_, err := mongoServiceOrderRepo.Delete(req.ID)

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

	return service_order.NewDeleteServiceOrderNoContent()
}

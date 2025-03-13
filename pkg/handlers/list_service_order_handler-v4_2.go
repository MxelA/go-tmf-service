package handler_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/repository"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func ListServiceOrderHandler(req service_order.ListServiceOrderParams) middleware.Responder {

	mongoServiceOrderRepo := repository.MongoServiceOrderRepository{
		MongoInstance: database.GetMongoInstance(),
		Context:       req.HTTPRequest.Context(),
	}

	queryParams := utils.BuildMongoFilter(req.HTTPRequest.URL.Query())
	retrieveServiceOrders, retrieveServiceOrdersTotalCount, err := mongoServiceOrderRepo.GetAllPaginate(queryParams, req.Fields, req.Offset, req.Limit)
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
		WithXTotalCount(*retrieveServiceOrdersTotalCount).
		WithXResultCount(int64(len(retrieveServiceOrders))).
		WithPayload(retrieveServiceOrders)
}

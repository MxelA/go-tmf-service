package handler_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (h *ServiceOrderHandler) ListServiceOrderHandler(req service_order.ListServiceOrderParams) middleware.Responder {

	queryParams := utils.BuildMongoFilter(req.HTTPRequest.URL.Query())
	retrieveServiceOrders, retrieveServiceOrdersTotalCount, err := h.repo.GetAllPaginate(req.HTTPRequest.Context(), queryParams, req.Fields, req.Offset, req.Limit)
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

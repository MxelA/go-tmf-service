package handler_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func (h *ServiceOrderHandler) RetrieveServiceOrderHandler(req service_order.RetrieveServiceOrderParams) middleware.Responder {

	retrieveServiceOrder, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, req.Fields)

	if err != nil {
		errCode := "500"
		reason := err.Error()
		var errModel = models.Error{
			Reason:  &reason,
			Code:    &errCode,
			Message: "Internal server error",
		}
		log.Println(err)
		return service_order.NewRetrieveServiceOrderInternalServerError().WithPayload(&errModel)
	}

	return service_order.NewRetrieveServiceOrderOK().WithPayload(retrieveServiceOrder)
}

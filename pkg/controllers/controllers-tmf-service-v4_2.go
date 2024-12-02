package controllers_tmf_service_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
)

func CreateServiceOrderControllerHandler(service_order.CreateServiceOrderParams) middleware.Responder {
	var r models.ServiceOrder
	return service_order.NewCreateServiceOrderCreated().WithPayload(&r)
}

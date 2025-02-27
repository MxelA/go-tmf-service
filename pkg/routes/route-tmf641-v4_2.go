package route_tmf641_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/handlers"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
)

var RegisterTmfServiceV4_2Routes = func(api *operations.ServiceOrderingV42API) {
	api.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(handler_v4_2.CreateServiceOrderHandler)
	api.ServiceOrderRetrieveServiceOrderHandler = service_order.RetrieveServiceOrderHandlerFunc(handler_v4_2.RetrieveServiceOrderHandler)
	api.ServiceOrderListServiceOrderHandler = service_order.ListServiceOrderHandlerFunc(handler_v4_2.ListServiceOrderHandler)
	api.ServiceOrderDeleteServiceOrderHandler = service_order.DeleteServiceOrderHandlerFunc(handler_v4_2.DeleteServiceOrderHandler)
}

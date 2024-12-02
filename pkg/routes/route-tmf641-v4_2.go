package route_tmf641_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
)

var RegisterTmfServiceV4_1Routes = func(api *operations.ServiceOrderingV42API) {
	api.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(CreateServiceOrderHandler)
}

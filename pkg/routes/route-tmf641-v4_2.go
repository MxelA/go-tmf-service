package route_tmf641_v4_2

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
	"github.com/MxelA/tmf-service-go/pkg/handlers"
	"github.com/MxelA/tmf-service-go/pkg/repository"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
)

var RegisterTmfServiceV4_2Routes = func(api *operations.ServiceOrderingV42API) {

	//Register dependency injection
	mongoInstance := database.GetMongoInstance()
	repo := &repository.MongoServiceOrderRepository{MongoInstance: mongoInstance}
	handler := handler_v4_2.NewServiceOrderHandler(repo)

	api.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(handler.CreateServiceOrderHandler)
	api.ServiceOrderRetrieveServiceOrderHandler = service_order.RetrieveServiceOrderHandlerFunc(handler.RetrieveServiceOrderHandler)
	api.ServiceOrderListServiceOrderHandler = service_order.ListServiceOrderHandlerFunc(handler.ListServiceOrderHandler)
	api.ServiceOrderDeleteServiceOrderHandler = service_order.DeleteServiceOrderHandlerFunc(handler.DeleteServiceOrderHandler)
	api.ServiceOrderPatchServiceOrderHandler = service_order.PatchServiceOrderHandlerFunc(handler.PatchServiceOrderHandler)
}

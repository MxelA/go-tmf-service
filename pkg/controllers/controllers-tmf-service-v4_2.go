package controllers_tmf_service_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/MxelA/tmf-service-go/pkg/utils"
	"github.com/go-openapi/runtime/middleware"
)

func CreateServiceOrderControllerHandler(req service_order.CreateServiceOrderParams) middleware.Responder {
	//dump, err := httputil.DumpRequestOut(req.HTTPRequest, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%q", dump)
	utils.PrettyPrint(req.ServiceOrder)
	//fmt.Printf("%+v\n", req.ServiceOrder)

	var res models.ServiceOrder
	return service_order.NewCreateServiceOrderCreated().WithPayload(&res)
}

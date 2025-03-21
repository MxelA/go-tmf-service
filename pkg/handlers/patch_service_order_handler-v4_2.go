package handler_v4_2

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"strings"
)

func (h *ServiceOrderHandler) PatchServiceOrderHandler(req service_order.PatchServiceOrderParams) middleware.Responder {
	contentType := req.HTTPRequest.Header["Content-Type"][0]

	if strings.Contains(contentType, "application/json-patch+json") {
		return processJsonPatch(h, req)
	} else if strings.Contains(contentType, "application/merge-patch+json") {
		return processMergePatch(h, req)
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

func processJsonPatch(h *ServiceOrderHandler, req service_order.PatchServiceOrderParams) middleware.Responder {
	//TO DO: Add JSON Patch support
	errCode := "422"
	reason := "Unsupported media type "
	var errModel = models.Error{
		Reason:  &reason,
		Code:    &errCode,
		Message: "Unsupported 'application/json-patch+json' media type ",
	}
	return service_order.NewPatchServiceOrderInternalServerError().WithPayload(&errModel)
}

func processMergePatch(h *ServiceOrderHandler, req service_order.PatchServiceOrderParams) middleware.Responder {

	_, err := h.repo.MergePatch(req.HTTPRequest.Context(), req.ID, req.ServiceOrder)
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

	retrieveServiceOrder, err := h.repo.GetByID(req.HTTPRequest.Context(), req.ID, nil)

	return service_order.NewPatchServiceOrderOK().WithPayload(retrieveServiceOrder)
}

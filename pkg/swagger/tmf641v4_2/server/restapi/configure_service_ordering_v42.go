// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/cancel_service_order"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/events_subscription"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/notification_listeners_client_side"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
)

//go:generate swagger generate server --target ../../server --name ServiceOrderingV42 --spec ../../TMF641-ServiceOrdering-v4.2.0.json --template-dir ./templates --principal interface{} --exclude-main

func configureFlags(api *operations.ServiceOrderingV42API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ServiceOrderingV42API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CancelServiceOrderCreateCancelServiceOrderHandler == nil {
		api.CancelServiceOrderCreateCancelServiceOrderHandler = cancel_service_order.CreateCancelServiceOrderHandlerFunc(func(params cancel_service_order.CreateCancelServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_service_order.CreateCancelServiceOrder has not yet been implemented")
		})
	}
	if api.ServiceOrderCreateServiceOrderHandler == nil {
		api.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(func(params service_order.CreateServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation service_order.CreateServiceOrder has not yet been implemented")
		})
	}
	if api.ServiceOrderDeleteServiceOrderHandler == nil {
		api.ServiceOrderDeleteServiceOrderHandler = service_order.DeleteServiceOrderHandlerFunc(func(params service_order.DeleteServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation service_order.DeleteServiceOrder has not yet been implemented")
		})
	}
	if api.CancelServiceOrderListCancelServiceOrderHandler == nil {
		api.CancelServiceOrderListCancelServiceOrderHandler = cancel_service_order.ListCancelServiceOrderHandlerFunc(func(params cancel_service_order.ListCancelServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_service_order.ListCancelServiceOrder has not yet been implemented")
		})
	}
	if api.ServiceOrderListServiceOrderHandler == nil {
		api.ServiceOrderListServiceOrderHandler = service_order.ListServiceOrderHandlerFunc(func(params service_order.ListServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation service_order.ListServiceOrder has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToCancelServiceOrderCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToCancelServiceOrderCreateEventHandler = notification_listeners_client_side.ListenToCancelServiceOrderCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToCancelServiceOrderCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToCancelServiceOrderCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToCancelServiceOrderInformationRequiredEventHandler == nil {
		api.NotificationListenersClientSideListenToCancelServiceOrderInformationRequiredEventHandler = notification_listeners_client_side.ListenToCancelServiceOrderInformationRequiredEventHandlerFunc(func(params notification_listeners_client_side.ListenToCancelServiceOrderInformationRequiredEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToCancelServiceOrderInformationRequiredEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToCancelServiceOrderStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToCancelServiceOrderStateChangeEventHandler = notification_listeners_client_side.ListenToCancelServiceOrderStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToCancelServiceOrderStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToCancelServiceOrderStateChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderAttributeValueChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderAttributeValueChangeEventHandler = notification_listeners_client_side.ListenToServiceOrderAttributeValueChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderAttributeValueChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderAttributeValueChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderCreateEventHandler = notification_listeners_client_side.ListenToServiceOrderCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderDeleteEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderDeleteEventHandler = notification_listeners_client_side.ListenToServiceOrderDeleteEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderDeleteEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderDeleteEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderInformationRequiredEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderInformationRequiredEventHandler = notification_listeners_client_side.ListenToServiceOrderInformationRequiredEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderInformationRequiredEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderInformationRequiredEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderJeopardyEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderJeopardyEventHandler = notification_listeners_client_side.ListenToServiceOrderJeopardyEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderJeopardyEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderJeopardyEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderMilestoneEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderMilestoneEventHandler = notification_listeners_client_side.ListenToServiceOrderMilestoneEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderMilestoneEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderMilestoneEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToServiceOrderStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToServiceOrderStateChangeEventHandler = notification_listeners_client_side.ListenToServiceOrderStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToServiceOrderStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToServiceOrderStateChangeEvent has not yet been implemented")
		})
	}
	if api.ServiceOrderPatchServiceOrderHandler == nil {
		api.ServiceOrderPatchServiceOrderHandler = service_order.PatchServiceOrderHandlerFunc(func(params service_order.PatchServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation service_order.PatchServiceOrder has not yet been implemented")
		})
	}
	if api.EventsSubscriptionRegisterListenerHandler == nil {
		api.EventsSubscriptionRegisterListenerHandler = events_subscription.RegisterListenerHandlerFunc(func(params events_subscription.RegisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.RegisterListener has not yet been implemented")
		})
	}
	if api.CancelServiceOrderRetrieveCancelServiceOrderHandler == nil {
		api.CancelServiceOrderRetrieveCancelServiceOrderHandler = cancel_service_order.RetrieveCancelServiceOrderHandlerFunc(func(params cancel_service_order.RetrieveCancelServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation cancel_service_order.RetrieveCancelServiceOrder has not yet been implemented")
		})
	}
	if api.ServiceOrderRetrieveServiceOrderHandler == nil {
		api.ServiceOrderRetrieveServiceOrderHandler = service_order.RetrieveServiceOrderHandlerFunc(func(params service_order.RetrieveServiceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation service_order.RetrieveServiceOrder has not yet been implemented")
		})
	}
	if api.EventsSubscriptionUnregisterListenerHandler == nil {
		api.EventsSubscriptionUnregisterListenerHandler = events_subscription.UnregisterListenerHandlerFunc(func(params events_subscription.UnregisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.UnregisterListener has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

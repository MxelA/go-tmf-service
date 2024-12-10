package main

import (
	"github.com/MxelA/tmf-service-go/pkg/routes"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"log"
)

func main() {

	var serverConfig restapi.Server

	// Parse flags
	parser := flags.NewParser(&serverConfig, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	//Connect to Mongo

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewServiceOrderingV42API(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.TLSCertificate = serverConfig.TLSCertificate
	server.TLSCertificateKey = serverConfig.TLSCertificateKey
	server.TLSPort = serverConfig.TLSPort

	route_tmf641_v4_2.RegisterTmfServiceV4_2Routes(api)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//})
	//
	//log.Println("Listening on localhost:8080")
	//
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

//Health route returns OK
//func CreateServiceOrderHandler(service_order.CreateServiceOrderParams) middleware.Responder {
//	var r models.ServiceOrder
//	return service_order.NewCreateServiceOrderCreated().WithPayload(&r)
//}

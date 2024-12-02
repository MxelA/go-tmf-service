package main

import (
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/models"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations/service_order"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	flags "github.com/jessevdk/go-flags"
	"log"
)

type Server struct {
	TLSCertificate flags.Filename `long:"tls-certificate" description:"The certificate file to use for secure connections" env:"TLS_CERTIFICATE" required:"true"`
	TLSKey         flags.Filename `long:"tls-key" description:"The private key to use for secure connections" env:"TLS_PRIVATE_KEY" required:"true"`
}

func main() {

	var serverConfig Server

	// Parse flags
	parser := flags.NewParser(&serverConfig, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

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

	server.Port = 8080
	server.TLSCertificate = serverConfig.TLSCertificate
	server.TLSCertificateKey = serverConfig.TLSKey

	api.ServiceOrderCreateServiceOrderHandler = service_order.CreateServiceOrderHandlerFunc(CreateServiceOrderHandler)

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
func CreateServiceOrderHandler(service_order.CreateServiceOrderParams) middleware.Responder {
	var r models.ServiceOrder
	return service_order.NewCreateServiceOrderCreated().WithPayload(&r)
}

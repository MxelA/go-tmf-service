package main

import (
	database "github.com/MxelA/tmf-service-go/pkg/config"
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
	if err := database.DbConnect(); err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
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

	server.TLSCertificate = serverConfig.TLSCertificate
	server.TLSCertificateKey = serverConfig.TLSCertificateKey
	server.TLSPort = serverConfig.TLSPort

	// Register routes
	route_tmf641_v4_2.RegisterTmfServiceV4_2Routes(api)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi"
	"github.com/MxelA/tmf-service-go/pkg/swagger/tmf641v4_2/server/restapi/operations"
	"github.com/go-openapi/loads"
	"html"
	"log"
	"net/http"
)

func main() {

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Listening on localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

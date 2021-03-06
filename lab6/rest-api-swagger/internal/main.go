package main

import (
	"log"

	"lab6/lab6/rest-api-swagger/pkg/swagger/server/restapi"
	"lab6/lab6/rest-api-swagger/pkg/swagger/server/restapi/operations"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			// error handle
			log.Fatalln(err)
		}
	}()

	server.Port = 8080

	// Implement the CheckHealth handler
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(
		func(user operations.CheckHealthParams) middleware.Responder {
			return operations.NewCheckHealthOK().WithPayload("OK\n")
		})

	// Implement the GetHelloUser handler
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(
		func(user operations.GetHelloUserParams) middleware.Responder {
			return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
		})

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

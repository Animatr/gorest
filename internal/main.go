package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Animatr/gorest/pkg/swagger/server/restapi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/Animatr/gorest/pkg/swagger/server/restapi/operations"
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

	api.CheckPriceHandler = operations.CheckPriceHandlerFunc(Price)

	api.GetProductProductHandler = operations.GetProductProductHandlerFunc(GetProductName)

	api.GetImageProductHandler = operations.GetImageProductHandlerFunc(GetImageByName)

	// Start server which listening
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

// Health route returns OK
func Price(operations.CheckPriceParams) middleware.Responder {
	return operations.NewCheckPriceOK().WithPayload("OK")
}

// GetHelloUser returns Hello + your name
func GetProductName(user operations.GetProductProductParams) middleware.Responder {
	return operations.NewGetProductProductOK().WithPayload("Hello " + user.Product + "!")
}

// GetGopherByName returns a gopher in png
func GetImageByName(image operations.GetImageProductParams) middleware.Responder {

	var URL string
	if image.Product != "" {
		URL = "https://github.com/scraly/gophers/raw/main/" + image.Product + ".png"
	} else {
		//by default we return dr who gopher
		URL = "https://github.com/scraly/gophers/raw/main/dr-who.png"
	}

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("error")
	}

	return operations.NewGetImageProductOK().WithPayload(response.Body)
}

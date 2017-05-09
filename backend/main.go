package app

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"

	"github.com/MiCHiLU/goapp-scaffold/restapi"
	"github.com/MiCHiLU/goapp-scaffold/restapi/operations"
)

func init() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewExampleAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.ConfigureFlags()
	server.ConfigureAPI()
	http.Handle("/", server.GetHandler())
	//http.HandleFunc("/", handleFunc)
}

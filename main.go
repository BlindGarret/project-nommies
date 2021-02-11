package main

import (
	"log"

	"github.com/BlindGarret/project-nommies/generated/api/restapi"
	"github.com/BlindGarret/project-nommies/generated/api/restapi/operations"
	"github.com/BlindGarret/project-nommies/handlers"
	"github.com/alexflint/go-arg"
	"github.com/go-openapi/loads"
)

type cliArgs struct {
	Port int `arg:"-p,help:port to listen to"`
}

var args = &cliArgs{
	Port: 2930,
}

func main() {
	arg.MustParse(args)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewRecipesAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = args.Port
	registerHandlers(api)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func registerHandlers(api *operations.RecipesAPI) {
	// Recipe Handlers
	api.RecipesListRecipesHandler = handlers.NewListRecipesHandler()
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"github.com/graphql-go/graphql"
	"../internal/utils"
)

func main() {
	fmt.Println("Launching the GraphQL server...")
	err := os.Chdir("../data")
	utils.Check(err)

	// Schema
	queryfields := graphql.Fields{
		"readFile": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"filename": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				filename, _ := params.Args["filename"].(string)

				return utils.Readdata(filename), nil
			},
		},	}

	mutationfields := graphql.Fields{
		"createNewFile": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"filename": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"text": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{} , error) {
				filename, _ := params.Args["filename"].(string)
				text, _ := params.Args["text"].(string)

				utils.Writedata(filename, text)
				return "ok", nil
			},
		},	
		"addToFile": &graphql.Field{
			Type: graphql.String,
			Args: graphql.FieldConfigArgument{
				"filename": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"text": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{} , error) {
				filename, _ := params.Args["filename"].(string)
				text, _ := params.Args["text"].(string)

				utils.Adddata(filename, text)
				return "ok", nil
			},
		},		
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: queryfields}
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields:	mutationfields,
	})
	schema, err := graphql.NewSchema(graphql.SchemaConfig {
		Query:		graphql.NewObject(rootQuery),
		Mutation:	rootMutation,
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// set up the graphql server on http port 8080
	http.HandleFunc("/graphql", func(write http.ResponseWriter, request *http.Request) {
		result := utils.ExecuteQuery(request.URL.Query().Get("query"), schema)
		json.NewEncoder(write).Encode(result)
	})

	fmt.Println("Now GraphQL Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"github.com/graphql-go/graphql"
)

func main() {
	fmt.Println("Launching the GraphQL server...")
	err := os.Chdir("../data")
	check(err)

	// Schema
	queryfields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return readdata("hello"), nil
			},
		},
		"test": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return readdata("test"), nil
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

				// save text to filename here
				writedata(filename, text)
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

				// save text to filename here
				adddata(filename, text)
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

	// Query
	query := `
		{
			hello
			test
		}		
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operations, errors: %+v", result.Errors)
	}
	resultJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", resultJSON) 

	// set up the graphql server on http port 8080
	http.HandleFunc("/graphql", func(write http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		json.NewEncoder(write).Encode(result)
	})

	fmt.Println("Now GraphQL Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

// To test, run GraphiQL and visit endpoint http://localhost:8080/graphql?
// with the following graphql query
// query basichttptest {
//	 hello
//	 test
// }
// To test graphql mutation
// mutation {
// 	createNewFile(
// 	  filename: "australia",
// 	  text: "hello from australia",
// 	)
//   }

// or using a browser with the following http request
// http://localhost:8080/graphql?query={hello}
// http://localhost:8080/graphql?query={test}
// or
// http://localhost:8080/graphql?query={hello%20test}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readdata(value string) string {
	f, err := os.Open(value)
	check(err)
	reader1, err := ioutil.ReadAll(f)
	check(err)
	f.Close()
	return string(reader1)
}

func writedata(filename string, text string) {
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	bytedata := []byte(text)
	_, err = f.Write(bytedata)
	check(err)
}

func adddata(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	fmt.Println("file named", filename, "opened")
	f.WriteString("\n"+text)
	check(err)
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:			schema,
		RequestString:	query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
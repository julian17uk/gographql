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

func readdata(value string) string {
	f, err := os.Open(value)
	check(err)
	reader1, err := ioutil.ReadAll(f)
	check(err)
	f.Close()
	return string(reader1)
}

func main() {
	fmt.Println("Launching the GraphQL server...")
	err := os.Chdir("../data")
	check(err)

	// Schema
	fields := graphql.Fields{
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

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			hello
		}		
	`

	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operations, errors: %+v", result.Errors)
	}
	resultJSON, _ := json.Marshal(result)
	fmt.Printf("%s \n", resultJSON) 

	check(err)
	h, err := os.Create("hello2")
	check(err)
	values := []byte("THis is a new string")
	h.Write(values)
	h.Close()
	i, err := os.Open("hello2")
	check(err)

	reader2, err := ioutil.ReadAll(i)
	i.Close()
	fmt.Printf("hello2 holds values: %s\n", string(reader2))

	fmt.Println("Finishing filepath access tests")
	// set up the graphql server on http port 8080
	http.HandleFunc("/graphql", func(write http.ResponseWriter, request *http.Request) {
		result := executeQuery(request.URL.Query().Get("query"), schema)
		json.NewEncoder(write).Encode(result)
	})

	fmt.Println("Now GraphQL Server is running on port 8080")
// To test, run GraphiQL and visit endpoint http://localhost:8080/graphql?
// with the following graphql query
// query basichttptest {
//	 hello
//	 test
// }

// or using a browser with the following http request
// http://localhost:8080/graphql?query={hello}
// http://localhost:8080/graphql?query={test}
// or
// http://localhost:8080/graphql?query={hello%20test}

	http.ListenAndServe(":8080", nil)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
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
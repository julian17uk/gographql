package main

import (
	"fmt"
	"testing"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"../internal/utils"
)

func TestWritedata(t *testing.T) {
	fmt.Println("running TestWritedata")
	filename := "Australia"
	text := "The capital is Canberra"
	expectedResult := "The capital is Canberra"

	utils.Writedata(filename, text)

	f, err := os.Open(filename)
	utils.Check(err)
	read, err := ioutil.ReadAll(f)
	utils.Check(err)
	f.Close()
	if read == nil {
		t.Errorf("Writedata Test returned nil")
	}
	if string(read) != expectedResult {
		t.Errorf("Writedata test failure")
	}

}

func TestAdddata(t *testing.T) {
	fmt.Println("running TestAdddata")
	filename := "Australia"
	text := "A major city is Sydney"
	expectedResult := "A major city is Sydney"

	utils.Adddata(filename, text)

	f, err := os.Open(filename)
	utils.Check(err)
	read, err := ioutil.ReadAll(f)
	utils.Check(err)
	f.Close()
	if read == nil {
		t.Errorf("Adddata test returned nil")
	}
	if !strings.Contains(string(read), expectedResult) {
		t.Errorf("Adddata test failure")
	}

}

func TestReaddata(t *testing.T) {
	fmt.Println("running TestReaddata")
	filename := "Australia"
	expectedResult := "The capital is Canberra\nA major city is Sydney"

	result := utils.Readdata(filename)

	if string(result)!=expectedResult {
		t.Errorf("Readdata test failure")
	}
}

func TestExecuteQuery(t *testing.T) {
	fmt.Println("running TestExecuteQuery")
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
	testquery := `
		{
			readFile(
				filename: "Australia"
			)
		}		
	`	
	
	expectedResult := "The capital is Canberra\nA major city is Sydney"

	result := utils.ExecuteQuery(testquery, schema)
	resultJSON, _ := json.Marshal(result)

	var g interface{}
	json.Unmarshal([]byte(resultJSON), &g)

	m := g.(map[string]interface{})
	datamap := m["data"]
	v := datamap.(map[string]interface{})
	dataoutput := v["readFile"].(string)

	fmt.Printf("JSON data: %s \n", dataoutput)

	if dataoutput != expectedResult {
		t.Errorf("ExecuteQuery test failure")
	}
}


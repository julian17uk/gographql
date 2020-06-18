package utils

import (
	"os"
	"io/ioutil"
	"fmt"
	"github.com/graphql-go/graphql"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Readdata(value string) string {
	f, err := os.Open(value)
	Check(err)
	reader1, err := ioutil.ReadAll(f)
	Check(err)
	f.Close()
	return string(reader1)
}

func Writedata(filename string, text string) {
	f, err := os.Create(filename)
	Check(err)
	defer f.Close()

	bytedata := []byte(text)
	_, err = f.Write(bytedata)
	Check(err)
}

func Adddata(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	Check(err)
	defer f.Close()
	f.WriteString("\n"+text)
	Check(err)
}

func ExecuteQuery(queryvalue string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:			schema,
		RequestString:	queryvalue,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
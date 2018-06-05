package main

import (
"net/http"
"github.com/graphql-go/graphql-go-handler"
"fmt"
)


func main() {

	initGraphQl()

	h := handler.New(&handler.Config{
		Schema: &pricecalculationSchema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8080", nil)
}
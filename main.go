package pricecalculator_go_graphql

import (
"net/http"
"github.com/graphql-go/graphql-go-handler"
"github.com/graphql-go/graphql"
"fmt"
)

type user struct {
	ID   string
	Name string
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOK := p.Args["id"].(string)
					if isOK {
						return doSomethingFancy(idQuery)
						//return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	})

//function that does something with the id the client send use
func doSomethingFancy(idQuery string) (user, error) {
	return user{ID:idQuery, Name:"Hello World"}, nil
}


var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

func main() {

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	// serve HTTP
	http.Handle("/graphql", h)

	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8080", nil)
}
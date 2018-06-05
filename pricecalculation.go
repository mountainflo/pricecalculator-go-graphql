package pricecalculator_go_graphql


import (
	"github.com/graphql-go/graphql"
)

type user struct {
	ID   string
	Name string
}

var (
	pricecalculationSchema graphql.Schema
)

func initGraphQl(){

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


	pricecalculationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}

//function that does something with the id the client send use
func doSomethingFancy(idQuery string) (user, error) {
	return user{ID:idQuery, Name:"Hello World"}, nil
}
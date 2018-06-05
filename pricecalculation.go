package main


import (
	"github.com/graphql-go/graphql"
)

type Item struct {
	ID   int
	Name string
	discout_perc float32
}

type CalcType int

const (
	RENTAL    CalcType = 0
	LEASING   CalcType = 1
)

var (
	pricecalculationSchema graphql.Schema
)

func initGraphQl(){

	/*
		for converting the input to an Item object

		input Item {
  			id: Int
  			name: String
		}
	 */
	var item = graphql.NewObject(graphql.ObjectConfig{
			Name: "item",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),	//NewNonNull if really necessary
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.ID, nil
						}
						return nil, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.Name, nil
						}
						return nil, nil
					},
				},
				"discout_perc": &graphql.Field{
					Type: graphql.Float,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.discout_perc, nil
						}
						return nil, nil
					},
				},
			},
		},
	)

	/*
		enum CalcType {
		  RENTAL
		  LEASING
		}
	 */
	calcTypeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "CalcType",
		Description: "The type for the calculation.",
		Values: graphql.EnumValueConfigMap{
			"RENTAL": &graphql.EnumValueConfig{
				Value:       RENTAL,
			},
			"LEASING": &graphql.EnumValueConfig{
				Value:       LEASING,
			},
		},
	})

	/*
		type RootQuery {
			calculate(type: CalcType, items [Item], family_discount: Int): [Item]
		}
	 */
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"calculate": &graphql.Field{
					Type: item,	//result of the request
					//Type: graphql.NewList(Item),	//result of the request
					Args: graphql.FieldConfigArgument{
						"type": &graphql.ArgumentConfig{
							Type: calcTypeEnum,
						},
						/*"item": &graphql.ArgumentConfig{
							Type: item,
							//Type: graphql.NewList(Item),
						},*/
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						//get the values from the request
						typeQuery, typeIsOK := p.Args["type"].(CalcType)
						//itemQuery, itemIsOK := p.Args["Item"].(Item)
						if typeIsOK {
							//do the calculation it input was correct
							return calculatePriceForItem(typeQuery)
						}
						return nil, nil
					},
				},
			},
		})

	//definition of the API
	pricecalculationSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
}

func calculatePriceForItem(typeQuery CalcType) (Item, error) {

	//TODO calculate price

	//return Item{ID:itemQuery.ID, Name:itemQuery.Name, discout_perc:itemQuery.discout_perc}, nil
	return Item{ID:1, Name:"asdf", discout_perc:1.5}, nil
}
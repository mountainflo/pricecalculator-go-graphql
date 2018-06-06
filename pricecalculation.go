package main


import (
	"github.com/graphql-go/graphql"
	"encoding/json"
)

type Item struct {
	Id   int	`json:"Id"`
	Name string	`json:"Name"`
	Discout_perc float32	`json:"Discout_perc"`
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
		for converting the output to an Item object
	 */
	var item = graphql.NewObject(graphql.ObjectConfig{
			Name: "item",
			Fields: graphql.Fields{
				"Id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),	//NewNonNull if really necessary
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.Id, nil
						}
						return nil, nil
					},
				},
				"Name": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.Name, nil
						}
						return nil, nil
					},
				},
				"Discout_perc": &graphql.Field{
					Type: graphql.Float,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.Discout_perc, nil
						}
						return nil, nil
					},
				},
			},
		},
	)


	/*
	 	converting input to an item
	 */
	var itemInput = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "itemInput",
			Fields: graphql.InputObjectConfigFieldMap{
				"Id": &graphql.InputObjectFieldConfig{
					Type: graphql.Int,
				},
				"Name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"Discout_perc": &graphql.InputObjectFieldConfig{
					Type: graphql.Float,
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
		Create the root query with:

		{
		  calculate( ... ) {
			...
		  }
		}

	 */
	var rootQuery = graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"calculate": &graphql.Field{
					Type: graphql.NewList(item),	//result-type for the response
					Args: graphql.FieldConfigArgument{
						"type": &graphql.ArgumentConfig{
							Type: calcTypeEnum,
						},
						"item": &graphql.ArgumentConfig{
							Type: graphql.NewList(itemInput),	//Input type of the request
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						//get the values from the request
						typeQuery, typeIsOK := p.Args["type"].(CalcType)


						//convert input to []Item{}-Object
						//not the best solution but works
						rJSON, _ := json.Marshal(p.Args["item"])
						//fmt.Printf("%s \n", rJSON)
						result := []Item{}
						err := json.Unmarshal(rJSON, &result)

						if typeIsOK && err == nil {
							//do the calculation it input was correct
							return calculatePriceForItem(typeQuery, result)
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

func calculatePriceForItem(typeQuery CalcType, itemQuery []Item) ([]Item, error) {

	//TODO calculate price

	//create the array result with all items
	result := []Item{{Id:itemQuery[0].Id, Name:itemQuery[0].Name, Discout_perc:itemQuery[0].Discout_perc}}

	//print("result.Name: ")
	//println(result[0].Name)

	return result, nil
}


/*
	TEST with

		{
		  calculate(type: RENTAL, item: [{Id: 3, Name: "blahNameConten", Discout_perc: 5.3}]) {
			Id
			Name
			Discout_perc
		  }
		}


 */
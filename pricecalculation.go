package main


import (
	"github.com/graphql-go/graphql"
	"encoding/json"
)

type Item struct {
	id   int
	name string
	discout_perc float32
}

type Blah struct {
	Id   string   `json:"id"`
	Name   string `json:"name"`
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
	/*var item = graphql.NewObject(graphql.ObjectConfig{
			Name: "item",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.Int),	//NewNonNull if really necessary
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.id, nil
						}
						return nil, nil
					},
				},
				"name": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						item, ok := p.Source.(Item)
						if ok {
							return item.name, nil
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
	)*/

	var blah = graphql.NewObject(graphql.ObjectConfig{
		Name: "blah",
		Fields: graphql.Fields{
			"Id": &graphql.Field{
				Type: graphql.String,	//NewNonNull if really necessary
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					blah, ok := p.Source.(Blah)
					if ok {
						return blah.Id, nil
					}
					return "ERROR", nil
				},
			},
			"Name": &graphql.Field{
				Type: graphql.String,	//NewNonNull if really necessary
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					blah, ok := p.Source.(Blah)
					if ok {
						return blah.Name, nil
					}
					return "ERROR", nil
				},
			},
		},
	},
	)

	var blahInput = graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "MyInputType",
			Fields: graphql.InputObjectConfigFieldMap{
				"Id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"Name": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
	)

	/*var blahInput = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "MyInputType",
			Fields: graphql.InputObjectConfigFieldMap{
				"key": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
			},
		},
		)*/

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
					Type: blah,	//result of the request
					//Type: graphql.NewList(Item),	//result of the request
					Args: graphql.FieldConfigArgument{
						"type": &graphql.ArgumentConfig{
							Type: calcTypeEnum,
						},
						"blah": &graphql.ArgumentConfig{
							Type: blahInput,
							//Type: graphql.NewList(Item),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						//get the values from the request
						typeQuery, typeIsOK := p.Args["type"].(CalcType)

						//print input of graphql
						rJSON, _ := json.Marshal(p.Args["blah"])
						result := Blah{}
						err := json.Unmarshal(rJSON, &result)
						print("result.Id: ")
						println(result.Id)

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

func calculatePriceForItem(typeQuery CalcType, itemQuery Blah) (interface{}, error) {

	//TODO calculate price

	//print(itemQuery)

	//return Item{ID:itemQuery.ID, Name:itemQuery.Name, discout_perc:itemQuery.discout_perc}, nil
	return Blah{Id:itemQuery.Id, Name:itemQuery.Name}, nil
}


/*
	TEST with

		{
		  calculate(type: RENTAL, blah: {Id: "blahContent", Name: "blahNameConten"}) {
			Id
			Name
		  }
		}

 */
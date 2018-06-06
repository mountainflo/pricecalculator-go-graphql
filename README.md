## Run local

Please note: You have to run the whole project not only the main function!


## Test with GraphiQL

* For testing use GraphiQL [https://electronjs.org/apps/graphiql](https://electronjs.org/apps/graphiql)
* GraphQL Endpoint: "http://localhost:8080/graphql"
* Method: GET

Request:

```
{
  calculate(type: RENTAL, item: [{Id: 3, Name: "FancySkiName", Discout_perc: 5.3}]) {
    Id
    Name
    Discout_perc
  }
}
```

If you for example only want ```id``` and ```name``` back, you can delete ```discout_perc``` from the request.

Response:
```
{
  "data": {
    "calculate": [
      {
        "Discout_perc": 5.3,
        "Id": 3,
        "Name": "blahNameConten"
      }
    ]
  }
}
```

## Current GraphQL-API
```
type RootQuery {
    calculate(type: CalcType, item: [Item]): [Item]
}

enum CalcType {
  RENTAL
  LEASING
}

input Item {
  Id: Int
  Name: string
  Discout_perc: float
}
```
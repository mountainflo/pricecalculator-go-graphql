## Run local

Please note: You have to run the whole project not only the main function!


## Test with GraphiQL

* For testing use GraphiQL [https://electronjs.org/apps/graphiql](https://electronjs.org/apps/graphiql)
* GraphQL Endpoint: "http://localhost:8080/graphql"
* Method: GET

Request:

```
{
  calculate(type: RENTAL) {
    id
    name
    discout_perc
  }
}
```

If you for example only want ```id``` and ```name``` back, you can delete ```discout_perc``` from the request.


Response:
```
{
  "data": {
    "calculate": {
      "discout_perc": 1.5,
      "id": 1,
      "name": "asdf"
    }
  }
}
```
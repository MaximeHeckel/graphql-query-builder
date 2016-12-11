# graphql-query-builder

A simple graphQL query builder built in Golang

## Getting started

To install the library run:

```
go get github.com/MaximeHeckel/graphql-query-builder/querybuilder
```

## Example

Here's a simple example which shows how to build a simple query using the library.

```
package main

import (
	"log"

	"github.com/MaximeHeckel/graphql-query-builder/querybuilder"
)

func main() {

	filter := querybuilder.NewQueryFilter(map[string]interface{}{"testfilter": "1234"})

	// NewQuery constructs the query
	// Filter is an optional argument to the constructor to set a parameter to 
	// run the query against
	query, _ := querybuilder.NewQuery("test", filter)
	query2, _ := querybuilder.NewQuery("nestedQuery")

	// Set an alias for the result of this query
	query2.SetAlias("myNestedQuery")

	// User find to set the properties to be returned from the query
	query2.Find([]interface{}{"nestedValue"})
	query.Find([]interface{}{"value", map[string]string{"hello": "test"}, map[string]int{"world": 2}, map[string]*querybuilder.Query{"testquery": query2}})

	// ToString() returns the string representation of the method
	log.Println(query.ToString())
}

```

## Tests

To run the tests:

```
go test -v
```
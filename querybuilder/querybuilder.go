package querybuilder

import (
	"errors"
	"fmt"
	"reflect"
)

// Query graphql query type
type Query struct {
	Alias  *string
	Name   string
	Filter map[string]interface{}
	Fields []interface{}
}

// QueryParams optional parameters of a query
type QueryParams struct {
	Filter map[string]interface{}
}

// ToString returns the actual query as a string
func (query *Query) ToString() string {
	var queryName string
	if query.Alias != nil {
		queryName = *query.Alias + ": " + query.Name
	} else {
		queryName = query.Name
	}

	if len(query.Filter) > 0 {
		for k, v := range query.Filter {
			queryName = queryName + "(" + fmt.Sprintf("%v:%#v", k, v) + ")"
		}
	}

	if len(query.Fields) > 0 {
		fields := ""
		for _, value := range query.Fields {
			switch reflect.TypeOf(value).String() {
			case "string":
				if fields != "" {
					fields = fields + "," + fmt.Sprintf("%+v", value)
				} else {
					fields = fields + fmt.Sprintf("%+v", value)
				}
			case "map[string]string":
				for k, v := range value.(map[string]string) {
					if fields != "" {
						fields = fields + "," + fmt.Sprintf("%v:%v", k, v)
					} else {
						fields = fields + fmt.Sprintf("%v:%v", k, v)
					}
				}
			case "map[string]int":
				for k, v := range value.(map[string]int) {
					if fields != "" {
						fields = fields + "," + fmt.Sprintf("%v:%v", k, v)
					} else {
						fields = fields + fmt.Sprintf("%v:%v", k, v)
					}
				}
			case "map[string]*querybuilder.Query":
				for k, v := range value.(map[string]*Query) {
					if fields != "" {
						fields = fields + "," + fmt.Sprintf("%v:%v", k, v.ToString())
					} else {
						fields = fields + fmt.Sprintf("%v:%v", k, v.ToString())
					}
				}
			}
		}
		queryName = queryName + "{" + fields + "}"
	}
	queryString := "{ " + queryName + " }"
	return queryString
}

// SetAlias sets an alias to the query
func (query *Query) SetAlias(alias string) *Query {
	query.Alias = &alias
	return query
}

// Find sets fields to the query
func (query *Query) Find(elements []interface{}) *Query {
	query.Fields = elements
	return query
}

// NewQueryFilter initiates an optional query filter
func NewQueryFilter(filter map[string]interface{}) func(*QueryParams) {
	optFilter := func(params *QueryParams) {
		params.Filter = filter
	}
	return optFilter
}

// NewQuery initiate a new query object
func NewQuery(name string, options ...func(*QueryParams)) (*Query, error) {
	if name == "" {
		return nil, errors.New("Invalid field name")
	}

	queryparams := QueryParams{}

	for _, param := range options {
		param(&queryparams)
	}

	if len(queryparams.Filter) > 1 {
		return nil, errors.New("Invalid field filter or alias")
	}
	newQuery := &Query{Name: name, Filter: queryparams.Filter}
	return newQuery, nil
}

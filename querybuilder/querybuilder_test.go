package querybuilder

import (
	"reflect"
	"testing"
)

func Test_ToString(t *testing.T) {
	testQueryToString := "{ test(testfilter:\"1234\"){value,hello:test,world:2,testquery:{ myNestedQuery: nestedQuery{nestedValue} }} }"

	filter := NewQueryFilter(map[string]interface{}{"testfilter": "1234"})
	query, _ := NewQuery("test", filter)
	query2, _ := NewQuery("nestedQuery")
	query2.SetAlias("myNestedQuery")
	query2.Find([]interface{}{"nestedValue"})
	query.Find([]interface{}{"value", map[string]string{"hello": "test"}, map[string]int{"world": 2}, map[string]*Query{"testquery": query2}})

	if query.ToString() != testQueryToString {
		t.Fatal("Expected string representation of queries to match")
	}
}

func Test_SetAlias(t *testing.T) {
	alias := "testAlias"
	testQuery := &Query{Name: "test", Alias: &alias}

	query, _ := NewQuery("test")
	query.SetAlias("testAlias")

	if !reflect.DeepEqual(query, testQuery) {
		t.Fatal("Expected queries to be equal")
	}
}

func Test_Find(t *testing.T) {
	testQuery := &Query{Name: "test", Fields: []interface{}{"value", map[string]string{"hello": "test"}, map[string]int{"world": 2}, map[string]*Query{"testquery": &Query{Name: "nestedQuery", Fields: []interface{}{"nestedValue"}}}}}

	query, _ := NewQuery("test")
	query2, _ := NewQuery("nestedQuery")
	query2.Find([]interface{}{"nestedValue"})
	query.Find([]interface{}{"value", map[string]string{"hello": "test"}, map[string]int{"world": 2}, map[string]*Query{"testquery": query2}})

	if !reflect.DeepEqual(query, testQuery) {
		t.Fatal("Expected queries to be equal")
	}
}

func Test_NewQuery(t *testing.T) {

	testQuery := &Query{Name: "test"}
	testQueryWithArgs := &Query{Name: "test2", Filter: map[string]interface{}{"testfilter": "1234"}}

	query, _ := NewQuery("test")
	myFilter := NewQueryFilter(map[string]interface{}{"testfilter": "1234"})
	query2, _ := NewQuery("test2", myFilter)

	myFilter2 := NewQueryFilter(map[string]interface{}{"testfilter": "1234", "testfilter2": "test"})
	query3, _ := NewQuery("test3", myFilter2)

	if !reflect.DeepEqual(query, testQuery) {
		t.Fatal("Expected queries to be equal")
	}

	if !reflect.DeepEqual(query2, testQueryWithArgs) {
		t.Fatal("Expected queries to be equal")
	}

	if query3 != nil {
		t.Fatal("Expected query to be nil")
	}
}

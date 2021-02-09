package querybuilder

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// QueryBuilder easy build query
func QueryBuilder(query string, entity interface{}) string {

	typeOf := reflect.TypeOf(entity)
	valueOf := reflect.ValueOf(entity)
	queryPointer := &query

	for index := 0; index < typeOf.NumField(); index++ {
		prepareQuery(typeOf, queryPointer, index)
		buildQuery(valueOf, queryPointer, index)
	}

	return query
}

func prepareQuery(typeOf reflect.Type, query *string, index int) {

	nameField := ":" + typeOf.Field(index).Name
	typeField := typeOf.Field(index).Type

	pattern := nameField + "(,)|" + nameField + "(,)?"
	re := regexp.MustCompile(pattern)

	strIdx := strconv.Itoa(index)
	if typeField.Kind() == reflect.String {
		strIdx = "'" + strIdx + "'"
	}
	strReplace := strIdx + "$1"

	replaced := re.ReplaceAll([]byte(*query), []byte(strReplace))
	*query = string(replaced)
}

func buildQuery(valueOf reflect.Value, query *string, index int) {
	var finalValue interface{}
	value := valueOf.Field(index)

	switch value.Kind() {
	case reflect.String:
		finalValue = value.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		finalValue = strconv.FormatInt(value.Int(), 10)
	default:
		fmt.Printf("unhandled kind %s", value.Kind())
	}
	strIdx := strconv.Itoa(index)
	*query = strings.Replace(*query, strIdx, fmt.Sprintf("%v", finalValue), -1)
}

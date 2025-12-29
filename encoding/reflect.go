package encoding

import (
	"fmt"
	"reflect"

	models "github.com/Lechenco/conversor-homebank/models/reflect"
)

func reflectFields(data interface{}) models.Elem {
	fields := []models.Field{}
	val := reflect.ValueOf(data)

	for i := 0; i < val.NumField(); i++ {

		if ignoreValue(val.Field(i)) {
			continue
		}

		name := val.Type().Field(i).Name
		tag := val.Type().Field(i).Tag.Get("qif")
		value := getValueAsString(val.Field(i))

		f := models.Field{
			Name:  models.Name(name),
			Tag:   models.Tag(tag),
			Value: models.Value(value),
		}
		fields = append(fields, f)
	}

	return models.Elem{Fields: fields}
}

func getValueAsString(v reflect.Value) string {
	value := v.String()

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		value = fmt.Sprintf("%.2f", v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = fmt.Sprintf("%d", v.Int())
	}

	return value
}

func ignoreValue(v reflect.Value) bool {

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Struct:
		return true
	default:
		return false
	}
}

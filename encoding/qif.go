package encoding

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/Lechenco/conversor-homebank/models/qif"
)

var nullableTags = []string{"M"}
var ignoreEmptyTags = []string{"L", "L[]"}

func Marshal(v interface{}) ([]byte, error) {
	output := ""
	var err error

	switch t := v.(type) {
	case qif.Account:
		output, err = marshalAccount(t)
	case []qif.Account:
		for _, acc := range t {
			o, _ := marshalAccount(acc)
			output = output + o
		}
	default:
		break
	}

	return []byte(output), err
}

func marshalAccount(v qif.Account) (string, error) {
	output := fmt.Sprintf("!Account\n%s\n", reflectTags(v))

	output = fmt.Sprintf("%s!Type:Bank\n", output)
	for _, t := range v.Transactions {
		output = fmt.Sprintf("%s%s\n", output, reflectTags(*t))
	}

	return output, nil
}

func reflectTags(data interface{}) string {
	output := ""
	val := reflect.ValueOf(data)
	for i := 0; i < val.NumField(); i++ {

		if ignoreReflect(val.Field(i)) {
			continue
		}

		tag := val.Type().Field(i).Tag.Get("qif")
		value := getValueString(val.Field(i))

		output = fmt.Sprintf("%s%s", output, formatValueForTag(value, tag))
	}

	return output + "^"
}

func getValueString(v reflect.Value) string {
	value := v.String()

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		value = fmt.Sprintf("%.2f", v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = fmt.Sprintf("%d", v.Int())
	}

	return value
}

func formatValueForTag(value, tag string) string {
	if value == "" && slices.Contains(nullableTags, tag) {
		return tag + "(null)\n"
	}
	if value == "" && slices.Contains(ignoreEmptyTags, tag) {
		return ""
	}

	if strings.Contains(tag, "[]") {
		return fmt.Sprintf("%c[%s]\n", tag[0], value)
	}

	return tag + value + "\n"
}

func ignoreReflect(v reflect.Value) bool {
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

func Unmarshal(data []byte, v interface{}) error {
	panic("")
}

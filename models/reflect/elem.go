package reflect

import (
	"fmt"
	"slices"
	"strings"
)

type Elem struct {
	Fields []Field
}

func (e Elem) Format(prefix string) string {
	output := prefix

	if prefix != "" {
		output += "\n"
	}

	for _, field := range e.Fields {
		format := formatValueForTag(field.Value, field.Tag)
		if format != "" {
			output += format + "\n"
		}
	}

	return output + "^"
}

func formatValueForTag(value Value, tag Tag) string {
	tags := string(tag)
	values := string(value)
	if values == "" && slices.Contains(nullableTags, tags) {
		return tags + "(null)"
	}
	if values == "" && slices.Contains(ignoreEmptyTags, tags) {
		return ""
	}

	if strings.Contains(tags, "[]") {
		return fmt.Sprintf("%c[%s]", tags[0], value)
	}

	return tags + values
}

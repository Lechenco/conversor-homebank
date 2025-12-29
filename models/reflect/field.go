package reflect

var nullableTags = []string{"M"}
var ignoreEmptyTags = []string{"L", "L[]"}

type Name string
type Tag string
type Value string

type Field struct {
	Name
	Tag
	Value
}

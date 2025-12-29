package encoding

import (
	"github.com/Lechenco/conversor-homebank/models"
)

func Marshal(v interface{}) ([]byte, error) {
	output := ""
	var err error

	switch t := v.(type) {
	case models.Account:
		output, err = marshalAccount(t)
	case []models.Account:
		for _, acc := range t {
			o, _ := marshalAccount(acc)
			output = output + o
		}
	default:
		break
	}

	return []byte(output), err
}

func marshalAccount(v models.Account) (string, error) {
	output := reflectFields(v).Format("!Account")

	output += "\n!Type:Bank\n"
	for _, t := range v.Transactions {
		output += reflectFields(*t).Format("") + "\n"
	}

	return output, nil
}

func Unmarshal(data []byte, v interface{}) error {
	panic("")
}

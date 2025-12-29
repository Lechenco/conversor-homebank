package io

import (
	"os"

	"github.com/Lechenco/conversor-homebank/encoding"
)

func WriteQIFFile(output string, data interface{}) {
	qifData, err := encoding.Marshal(data)

	if err != nil {
		panic("Error while parsing objects to QIF file: " + err.Error())
	}

	os.WriteFile(output, qifData, 0644)
}

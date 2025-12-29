package main

import (
	"flag"

	"github.com/Lechenco/conversor-homebank/io"
	"github.com/Lechenco/conversor-homebank/services"
)

func main() {
	input := flag.String("f", "data/data.csv", "Input CSV file")
	output := flag.String("o", "output.qif", "Output QIF file")
	flag.Parse()

	records := io.ReadCSV(*input)

	accounts := services.RecordsToAccounts(records)

	io.WriteQIFFile(*output, accounts)
}

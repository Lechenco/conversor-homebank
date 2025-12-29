package main

import (
	"fmt"
	"os"

	"github.com/Lechenco/conversor-homebank/encoding"
	"github.com/Lechenco/conversor-homebank/models/qif"
)

func main() {

	data := qif.Account{
		Name: "Teste",
		Transactions: []*qif.Transaction{
			{
				Date:     "2025/04/05",
				Value:    -180.98,
				Status:   qif.Reconcilied,
				Payee:    "Nicole",
				Memo:     "Transação de teste",
				Category: "Lazer",
			},
		},
	}

	str, err := encoding.Marshal([]qif.Account{data, data})
	fmt.Print(string(str))
	fmt.Print(err)

	os.WriteFile("data/output.qif", str, 0644)
}

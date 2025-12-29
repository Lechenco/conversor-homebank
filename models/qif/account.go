package qif

type Account struct {
	Name         string `qif:"N"`
	Transactions []*Transaction
}

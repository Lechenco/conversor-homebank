package models

type Transaction struct {
	Date     string  `qif:"D"`
	Value    float32 `qif:"T"`
	Status   Status  `qif:"C"`
	Payee    string  `qif:"P"`
	Memo     string  `qif:"M"`
	Category string  `qif:"L"`
	Transfer string  `qif:"L[]"`
}

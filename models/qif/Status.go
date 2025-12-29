package qif

type Status string

const (
	None        Status = ""
	Cleared     Status = "c"
	Reconcilied Status = "R"
)

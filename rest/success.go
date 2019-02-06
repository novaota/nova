package rest

type success struct {
	Operation string
	Success   bool
}

func NewSuccess(operation string) *success {
	return &success{Operation: operation, Success: true}
}


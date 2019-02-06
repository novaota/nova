package rest

type Error interface {
	error
	Status() int
}

type RestError struct {
	Err        error
	HttpStatus int
	Message    string
}

func (err RestError) Error() string {
	return err.Message
}

func NewRestError(status int) RestError {
	return RestError{HttpStatus: status}
}

func (err RestError) Status() int {
	return err.HttpStatus
}

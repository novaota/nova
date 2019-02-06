package updatepackage

type SignService interface {
	SignFile(file string) (string, error)
	IsSigned(signature string, file string) (bool, error)
}

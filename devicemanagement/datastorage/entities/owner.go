package entities

type Owner struct {
	GormEntity
	Name      string
	FirstName string
	Address   string
	Zip       string
}

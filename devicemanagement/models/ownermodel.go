package models

type OwnerModel struct {
	ModelBase
	Name      string
	FirstName string
	Address   string
	Zip       string
}

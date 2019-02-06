package models

type Model interface {
	Validate() error
}

type ModelBase struct {
	ID uint
}
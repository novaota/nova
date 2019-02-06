package entities

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

type Entity interface {
	Validate() error
	GetID() uint
}

//cached singleton instance
var validate *validator.Validate = validator.New()

type GormEntity struct {
	gorm.Model
}

func (model *GormEntity) Validate() error {
	return validate.Struct(model)
}

func (model *GormEntity) GetID() uint {
	return model.ID
}

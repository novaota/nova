package datastorage

import (
	"nova/devicemanagement/datastorage/entities"
)

type Repository interface {
	GetAll(out interface{})
	Get(id int, out interface{})
	AddOrUpdate(model entities.Entity) error
	Add(model interface{}) error
	Update(model entities.Entity) error
	First(out interface{}, predicate interface{})
	GetWhere(out interface{}, predicate interface{})
	Delete(model entities.Entity)
	Close()
	CreateSchema()
	IsSchemaCreated() bool
}

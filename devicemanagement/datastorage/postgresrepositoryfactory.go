package datastorage

import (
	// needed for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgresRepositoryFactory struct {
	ConnectionString string
}

func NewPostgresRepository(connectionString string) *postgresRepositoryFactory {
	return &postgresRepositoryFactory{connectionString}
}

func (factory *postgresRepositoryFactory) CreateRepository() Repository {
	return NewGormRepository("postgres", factory.ConnectionString)
}

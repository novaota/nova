package datastorage

type RepositoryFactory interface {
	CreateRepository() Repository
}
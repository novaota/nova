package implementation

import (
	"golang.org/x/crypto/bcrypt"
	"nova/devicemanagement/datastorage"
	"nova/devicemanagement/datastorage/entities"
)

type databaseAuthentificationService struct {
	datastorage.RepositoryFactory
}

func NewDatabaseAuthenticationService(repositoryFactory datastorage.RepositoryFactory) *databaseAuthentificationService {
	return &databaseAuthentificationService{RepositoryFactory: repositoryFactory}
}

func (service databaseAuthentificationService) RegisterUser(username string, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := &entities.User{Username: username, Password: hash}
	repo := service.RepositoryFactory.CreateRepository()
	defer repo.Close()

	repo.Add(user)
	return nil
}

func (service databaseAuthentificationService) IsRegistered(username string, password string) bool {
	expectedUser := &entities.User{Username: username}

	repo := service.RepositoryFactory.CreateRepository()
	defer repo.Close()

	foundUser := &entities.User{}
	repo.First(&foundUser, expectedUser)

	err := bcrypt.CompareHashAndPassword(foundUser.Password, []byte(password))
	return err == nil
}

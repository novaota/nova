package services

type AuthenticationService interface {
	RegisterUser(username string, password string) error
	IsRegistered(username string, password string) bool
}
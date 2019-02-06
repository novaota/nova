package models

type UserModel struct {
	ModelBase
	Username string
	Password string
	Role     string
}
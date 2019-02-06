package entities

type User struct {
	GormEntity
	Username string
	Password []byte
	Role     string
}

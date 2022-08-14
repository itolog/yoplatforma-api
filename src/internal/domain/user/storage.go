package user

type Storage interface {
	GetOne(id string) *User
}

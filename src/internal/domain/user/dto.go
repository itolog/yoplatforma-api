package user

type CreateUserDto struct {
	Name     string `validate:"required,min=3,max=15"`
	Password string `validate:"required,min=6,max=32"`
}

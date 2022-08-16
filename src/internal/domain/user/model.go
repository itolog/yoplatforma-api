package user

type Avatar struct {
	Url string
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Avatar   Avatar `json:"avatar"`
	Password string `json:"password"`
}

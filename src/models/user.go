package models

type User struct {
	ID       int
	Name     string
	Login    string
	password string
}

func (u *User) SetPassword(password string) {
	u.password = password
}

package models

type User struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	password string
}

func (u *User) SetPassword(password string) {
	u.password = password
}

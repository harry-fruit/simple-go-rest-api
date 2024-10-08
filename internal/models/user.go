package models

type User struct {
	ID       int    `json:"id" db:"id"`
	IdStatus int    `json:"id_status" db:"id_status"`
	IdRole   int    `json:"id_role" db:"id_role"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	password string
}

func (u *User) SetPassword(password string) {
	u.password = password
}

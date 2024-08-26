package dtos

type UserPayloadDTO struct {
	IdStatus int    `json:"id_status" db:"id_status"`
	IdRole   int    `json:"id_role" db:"id_role"`
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
}

type SetUserPasswordPayloadDTO struct {
	Password string `json:"password" db:"password"`
}

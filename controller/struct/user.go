package _struct

type UserDelete struct {
	UUID string `json:"uuid" binding:"required"`
}

type ResultUser struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
}

type UpdateUser struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

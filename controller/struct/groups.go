package _struct

type GroupsDelete struct {
	ID uint `json:"id" binding:"required"`
}

type GroupsInsert struct {
	Name     string   `json:"name" binding:"required"`
	Endpoind []uint   `json:"endpoind_id"`
	UserUUID []string `json:"user_uuid"`
}

type GroupsEndpoindInsert struct {
	ID       uint   `json:"id" binding:"required"`
	Endpoind []uint `json:"endpoind" binding:"required"`
}

type UserGroup struct {
	UUID_User string `json:"uuid_user" binding:"required"`
	GroupsID  []uint `json:"groups_id"`
}

package _struct

type PermisisonUserDelete struct {
	UUID string `json:"uuid" binding:"required"`
}

type PermissionUserInsert struct {
	UUID     string `json:"user_uuid" binding:"required"`
	Endpoind []uint `json:"endpoind_id" binding:"required"`
}

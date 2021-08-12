package model

import (
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	Token         string `json:"token,omitempty"`
	UserModelUUID string `json:"user_model_uuid,omitempty"`
	Status        bool   `gorm:"default:true" json:"status,omitempty"`
}

func (token Token) BeforeCreate(tx *gorm.DB) (err error) {
	deleteOldData(token.UserModelUUID)
	return
}

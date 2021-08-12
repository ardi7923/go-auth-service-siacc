package model

import (
	"auth/config"
	"auth/helper"
	"fmt"

	"gorm.io/gorm"
)

// => CRUD FUNCTION

// =========================> SERVICES <=========================

type ServicesResponse struct {
	ID           uint               `json:"id"`
	Name         string             `json:"service_name"`
	Url_Services string             `json:"url_services"`
	Endpoind     []ServicesEndpoind `json:"endpoind"`
}

func (serv ServicesModel) GetAll() (interface{}, error) {
	var response []ServicesModel
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Preload("Endpoind").Find(&response).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(&response, &[]ServicesResponse{}), nil
}

func (serv ServicesModel) Insert() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Create(&serv).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(serv, &ServicesResponse{}), nil
}

func (serv ServicesModel) Update() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Model(ServicesModel{}).Where("id = ? ", serv.ID).Updates(serv).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(serv, &ServicesResponse{}), nil
}

func (serv ServicesModel) Delete(id uint) (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	var endpoind []ServicesEndpoind
	var id_endpoind []uint

	if err := conn.Where("id = ? ", id).Delete(ServicesModel{}).Error; err != nil {
		return nil, err
	}

	conn.Where("services_model_id = ? ", id).Find(&endpoind)

	for _, value := range endpoind {
		id_endpoind = append(id_endpoind, value.ID)
	}

	conn.Where("id IN ( ? ) ", id_endpoind).Unscoped().Delete(&endpoind)
	conn.Exec("DELETE FROM user_endpoind_permission WHERE services_endpoind_id IN (?) ", id_endpoind)
	conn.Exec("DELETE FROM group_endpoind_permission WHERE services_endpoind_id IN (?) ", id_endpoind)
	return helper.SaveJsonResponse(serv, &ServicesResponse{}), nil
}

// =======================> END SERVICE <==========================

// =========================> ENDPOIND <===========================

func (endpoind *ServicesEndpoind) TableName() string { return "endpoind_services" }

type EndpoindResponse struct {
	ID              uint   `json:"id"`
	ServicesModelID uint   `json:"services_id" gorm:"index"`
	Name            string `json:"endpoind" binding:"required"`
	Method          string `json:"method" binding:"required"`
}

func (endpoind *ServicesEndpoind) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("delete enpoind")
	fmt.Println("delete after")
	return
}

func (end ServicesEndpoind) Insert() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Create(&end).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(end, &EndpoindResponse{}), nil
}

func (end *ServicesEndpoind) Update() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("id = ? ", end.ID).Save(&end).Error; err != nil {
		return nil, err
	}
	return end, nil
}

// =====================> END ENDPOIND <========================

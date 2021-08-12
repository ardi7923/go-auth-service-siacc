package model

import (
	"auth/config"
	"auth/helper"
	"errors"
)

func (user UserModel) TokenCheck(uuid string) (status bool, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("uuid = ? ", uuid).Find(&user).Error; err != nil {
		return false, err
	}
	if user.Status == false {
		return false, errors.New("User Status Disable")
	}
	if user.Token.Status == false {
		return false, errors.New("Token Status Disable")
	}
	return true, nil
}

func (user *UserModel) TokenIsActive() bool {
	return user.Token.Status
}

// s

// => CRUD FUNCTION

func (user UserModel) GetAll() (interface{}, error) {
	conn, db := config.DataBase()
	defer db.Close()
	var query = "SELECT * FROM " + user.TableName()
	var response []userModelResponse
	if err := conn.Raw(query).Scan(&response).Error; err != nil {
		return nil, err
	}
	return response, nil
}

func (user *UserModel) Insert() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Create(&user).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(user, &userModelResponse{}), nil
}

func (user UserModel) Update(uuid string, data_json UserModel) (interface{}, error) {
	conn, db := config.DataBase()
	defer db.Close()
	query := conn.Where("uuid = ? ", uuid).First(&user)
	user.Username = data_json.Username
	user.Password = data_json.Password
	if err := query.Error; err != nil {
		return nil, err
	}
	if err := conn.Save(&user).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(user, &userModelResponse{}), nil
}

func (user UserModel) Delete(uuid string) (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("uuid = ? ", uuid).Find(&user).Unscoped().Delete(&user).Error; err != nil {
		return nil, err
	}
	return helper.SaveJsonResponse(user, &userModelResponse{}), nil
}

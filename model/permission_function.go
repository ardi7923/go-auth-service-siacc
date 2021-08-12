package model

import (
	"auth/config"
	_struct "auth/controller/struct"
	"fmt"
)

// => HOOKS Function

func (user_permission UserPermissionModel) TableName() string { return "user_permission" }
func (groups GroupPermissionModel) TableName() string         { return "groups" }

//======== user_function permissions ===========//

func (permissionUser UserPermissionModel) UserPermission(uuid string, services string, endpoind string, method string) (*_struct.ResultUser, error) {
	conn, db := config.DataBase()
	defer db.Close()
	var user_result _struct.ResultUser

	var query = "SELECT * FROM " + UserModel{}.TableName() +
		" INNER JOIN user_permission ON user_permission.user_model_uuid = user_auth.uuid " +
		"INNER JOIN user_endpoind_permission ON user_endpoind_permission.user_permission_model_id = user_permission.id " +
		"INNER JOIN endpoind_services ON endpoind_services.id = user_endpoind_permission.services_endpoind_id " +
		"INNER JOIN list_services ON list_services.id = endpoind_services.services_model_id " +
		"WHERE user_auth.uuid = ? AND list_services.name = ? AND endpoind_services.name = ? AND endpoind_services.method = ?"

	if err := conn.Raw(query, uuid, services, endpoind, method).Scan(&user_result).Error; err != nil {
		fmt.Println(err)
		return &_struct.ResultUser{}, err
	}
	return &user_result, nil
}

func (permissionUser UserPermissionModel) Insert(data *_struct.PermissionUserInsert) (data_interface interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	var data_endpoind []ServicesEndpoind

	fmt.Println(data.UUID)
	fmt.Println(data.Endpoind)

	if err := conn.Where("id IN (?)", data.Endpoind).Find(&data_endpoind).Error; err != nil {
		return nil, err
	}
	permissionUser = UserPermissionModel{
		UserModelUUID:      data.UUID,
		ServicesEndpoindID: data_endpoind,
	}

	if err := conn.Create(&permissionUser).Error; err != nil {
		return nil, err
	}

	return permissionUser, nil
}

func (permission_user *UserPermissionModel) Update() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("user_model_uuid = ? ", permission_user.UserModelUUID).Save(&permission_user).Error; err != nil {
		return nil, err
	}
	return permission_user, nil
}

func (permission_user UserPermissionModel) Delete(uuid string) (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("user_model_uuid = ? ", uuid).Find(&permission_user).Unscoped().Delete(&permission_user).Error; err != nil {
		return nil, err
	}
	if err := conn.Exec("DELETE FROM user_endpoind_permission WHERE user_permission_model_id = ? ", permission_user.ID).Error; err != nil {
		fmt.Println(err)
	}
	return permission_user, nil
}

//=============== groups ==========//

func (groups GroupPermissionModel) GroupsPermission(uuid string, services string, endpoind string, method string) (*_struct.ResultUser, error) {
	conn, db := config.DataBase()
	defer db.Close()
	var user_result _struct.ResultUser

	const query = "SELECT * FROM user_auth " +
		"INNER JOIN user_groups ON user_groups.user_model_id = user_auth.id " +
		"INNER JOIN groups ON groups.id = user_groups.group_permission_model_id " +
		"INNER JOIN group_endpoind_permission ON group_endpoind_permission.group_permission_model_id = groups.id " +
		"INNER JOIN endpoind_services ON group_endpoind_permission.services_endpoind_id = endpoind_services.id " +
		"INNER JOIN list_services ON list_services.id = endpoind_services.services_model_id " +
		"WHERE user_auth.uuid = ? AND list_services.name = ? AND endpoind_services.name = ? AND endpoind_services.method = ?"

	if err := conn.Raw(query, uuid, services, endpoind, method).Scan(&user_result).Error; err != nil {
		return &_struct.ResultUser{}, err
	}
	return &user_result, nil
}

func (groups GroupPermissionModel) Delete(id uint) (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("id = ? ", id).Find(&groups).Association("EndpoindID").Clear(); err != nil {
		return nil, err
	}
	if err := conn.Model(&groups).Unscoped().Delete(&groups).Error; err != nil {
		return nil, err
	}
	conn.Exec("DELETE FROM user_groups WHERE group_permission_model_id = ? ", id)
	return groups, nil
}

func (groups *GroupPermissionModel) Update() (data interface{}, err error) {
	conn, db := config.DataBase()
	defer db.Close()
	if err := conn.Where("id = ? ", groups.ID).Save(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

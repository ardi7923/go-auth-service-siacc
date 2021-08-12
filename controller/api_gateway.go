package controller

import (
	"auth/config"
	_struct "auth/controller/struct"
	"auth/helper"
	"auth/model"
	"auth/response"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckPermissionUser(ctx *gin.Context) {
	var data_json _struct.ApiGateway
	var user_permission_model model.UserPermissionModel
	var group_permission_model model.GroupPermissionModel
	if err := ctx.ShouldBindJSON(&data_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": err.Error(),
			})
		return
	}
	claims, err_claims := helper.GetPayloadToken(data_json.Token)
	if err_claims != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"data":    nil,
				"message": "error",
			})
		return
	}
	uuid := claims["UUID"].(string)

	if user, err := user_permission_model.UserPermission(uuid, data_json.ServiceName, data_json.Endpoind, data_json.Method); err == nil {
		res, err1 := getUserByUUID(uuid)
		if err != nil {
			ctx.SecureJSON(response.BadRequest("error get permission"))
			return
		}
		if err1 != nil {
			ctx.SecureJSON(response.BadRequest("error get user detail"))
			return
		}

		if user.UUID != "" && user.Username != "" {
			ctx.SecureJSON(
				http.StatusOK,
				gin.H{
					"data": gin.H{
						"user":    user,
						"biodata": res.Data,
					},
					"message": "user_function have permission",
				})
			return
		}
	}

	if user, err := group_permission_model.GroupsPermission(uuid, data_json.ServiceName, data_json.Endpoind, data_json.Method); err == nil {
		res, err1 := getUserByUUID(uuid)
		if err != nil {
			ctx.SecureJSON(response.BadRequest("error get group permission"))
			return
		}
		if err1 != nil {
			ctx.SecureJSON(response.BadRequest("error get user detail"))
			return
		}
		if user.UUID != "" && user.Username != "" {
			ctx.SecureJSON(
				http.StatusOK,
				gin.H{
					"data": gin.H{
						"user":    user,
						"biodata": res.Data,
					},
					"message": "user_function have permission group",
				})
			return
		}

	}

	ctx.SecureJSON(
		http.StatusUnauthorized,
		gin.H{
			"data":    nil,
			"message": "rejected",
		})
}

// This function for get biodata user by uuid, make request to employee-biodata service
func getUserByUUID(uuid string) (_struct.Result, error) {
	var (
		payload = strings.NewReader(`{"uuid":"` + uuid + `"}`)
		token   = config.Token_service
		url     = config.Link_api_services + "/employee-biodata/employee/by-uuid"
	)

	res, err := helper.RequestUrl("POST", url, token, payload)

	defer res.Body.Close()

	if err != nil {
		return _struct.Result{}, err
	}

	response := _struct.Result{}
	jsonErr := json.NewDecoder(res.Body).Decode(&response)

	if jsonErr != nil {
		return _struct.Result{}, jsonErr
	}

	if res.StatusCode != 200 {
		return _struct.Result{}, errors.New(response.Message)
	}

	return response, nil
}

package controller

import (
	"auth/config"
	_struct "auth/controller/struct"
	"auth/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServicesGetAll(ctx *gin.Context) {
	var data_json model.ServicesModel
	response, err := data_json.GetAll()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
				"data":    nil,
			})
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data":    response,
		})
}

func ServicesInsert(ctx *gin.Context) {
	var service model.ServicesModel
	if err := ctx.ShouldBindJSON(&service); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := service.Insert()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": "success",
		})
	return
}

func EndpoindInsert(ctx *gin.Context) {
	conn, db := config.DataBase()
	defer db.Close()
	var enpoind model.ServicesEndpoind
	var service model.ServicesModel

	if err := ctx.ShouldBindJSON(&enpoind); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	if err := conn.Where("id = ?", enpoind.ServicesModelID).First(&service).Error; err != nil {
		ctx.SecureJSON(
			http.StatusNotFound,
			gin.H{
				"message": "not found",
				"data":    nil,
			})
		return
	}
	response, err := enpoind.Insert()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
				"data":    nil,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "status ok",
			"data":    response,
		})
}

func ServiceUpdate(ctx *gin.Context) {

	var services model.ServicesModel
	//var struct_model model.ServicesModel
	//var query 			*gorm.DB

	if err := ctx.ShouldBindJSON(&services); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
				"data":    nil,
			})
		return
	}
	response, err := services.Update()
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
				"data":    nil,
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "success",
			"data":    response,
		})
	//query = config.DataBase().Model(&data_json).Where("id = ?",data_json.ID).First(&struct_model)
	//data_json.ID = struct_model.ID
	//utils.ModelUpdate(ctx,&data_json,query)
}

func ServiceDelete(ctx *gin.Context) {
	var struct_json _struct.DeleteService
	if err := ctx.ShouldBindJSON(&struct_json); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	response, err := model.ServicesModel{}.Delete(struct_json.ID)
	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}
	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"message": "success to delete",
			"data":    response,
		})
	return

}

func EndpointUpdate(ctx *gin.Context) {

	var endpoind model.ServicesEndpoind

	if err := ctx.ShouldBindJSON(&endpoind); err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	response, err := endpoind.Update()

	if err != nil {
		ctx.SecureJSON(
			http.StatusNotAcceptable,
			gin.H{
				"message": err.Error(),
			})
		return
	}

	ctx.SecureJSON(
		http.StatusOK,
		gin.H{
			"data":    response,
			"message": err.Error(),
		})

	// query = config.DataBase().Where("id = ? and services_model_id = ?", data_json.ID, data_json.ServicesModelID)
	// data_json.ID = struct_model.ID
	// utils.ModelUpdate(ctx, &data_json, query)
}

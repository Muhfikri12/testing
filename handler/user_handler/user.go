package userhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewUserHandler(service service.AllService, log *zap.Logger, config util.Configuration) UserHandler {
	return UserHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (uh *UserHandler) GetListAddress(c *gin.Context) {

	token := c.GetHeader("Authorization")

	Addresses, err := uh.Service.UserService.GetListAddress(token)
	if err != nil {
		uh.Log.Error("Data not found: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Data not found", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully Get Data", Addresses)
}

func (uh *UserHandler) GetDetailUser(c *gin.Context) {

	token := c.GetHeader("Authorization")

	users, err := uh.Service.UserService.GetDetailUser(token)
	if err != nil {
		uh.Log.Error("Data not found: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Data not found", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully Get Data", users)
}

func (uh *UserHandler) UpdateUserData(c *gin.Context) {

	User := model.Users{}

	validate := validator.New()
	token := c.GetHeader("Authorization")

	err := c.ShouldBindJSON(&User)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.ResponsesJson(c, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(User)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(User)
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.UpdateUserData(token, &User)
	if err != nil {
		uh.Log.Error("Failed to update Data: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to update Data: "+err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Update Data", User)

}

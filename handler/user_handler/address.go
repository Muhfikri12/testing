package userhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (uh *UserHandler) AddAddress(c *gin.Context) {

	validate := validator.New()
	token := c.GetHeader("Authorization")
	address := model.Addresses{}

	err := c.ShouldBindJSON(&address)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.ResponsesJson(c, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(address)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(address)
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.AddAddress(token, &address)
	if err != nil {
		uh.Log.Error("Failed to create address: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to create address", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusCreated, "Succesfully Create address", address)
}

func (uh *UserHandler) UpdateAddress(c *gin.Context) {

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	validate := validator.New()
	token := c.GetHeader("Authorization")
	address := model.Addresses{}

	err := c.ShouldBindJSON(&address)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.ResponsesJson(c, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(address)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(address)
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.UpdateAddress(token, id, &address)
	if err != nil {
		uh.Log.Error("Failed to update address: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to update address", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully Update address", address)
}

func (uh *UserHandler) DeleteAddress(c *gin.Context) {

	token := c.GetHeader("Authorization")

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	err := uh.Service.UserService.DeleteAddress(token, id)
	if err != nil {
		uh.Log.Error("Failed to delete address: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully delete address", nil)
}

func (uh *UserHandler) SetDefault(c *gin.Context) {

	token := c.GetHeader("Authorization")

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	err := uh.Service.UserService.SetDefault(token, id)
	if err != nil {
		uh.Log.Error("Failed to Set Default address: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to Set Default address", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Succesfully Set Default address", nil)
}

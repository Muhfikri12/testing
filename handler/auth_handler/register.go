package authhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (u *AuthHandler) Register(c *gin.Context) {
	user := model.Register{}
	validate := validator.New()

	if err := c.ShouldBindJSON(&user); err != nil {
		u.Log.Error("Error from Register Handler :" + err.Error())
		helper.ResponsesJson(c, http.StatusBadRequest, "invalid request payload", nil)
		return
	}

	err := validate.Struct(user)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(user)
		helper.ResponsesJson(c, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	if err := u.Service.AuthService.Register(&user); err != nil {
		u.Log.Error("Failed to Register: " + err.Error())
		helper.ResponsesJson(c, http.StatusInternalServerError, "failed to Register: "+err.Error(), nil)

		return
	}

	helper.ResponsesJson(c, http.StatusCreated, "Successfully Register", user)

}

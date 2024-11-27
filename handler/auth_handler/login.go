package authhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewAuthHandler(service service.AllService, log *zap.Logger, config util.Configuration) AuthHandler {
	return AuthHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (l *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	user := model.Login{}
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		l.Log.Error("Invalid request payload: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid request payload: "+err.Error(), nil)
		return
	}

	err := validate.Struct(user)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(user)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	if err := l.Service.AuthService.Login(&user); err != nil {
		l.Log.Error("Failed to login: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to login: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Login ", user)
}

func (l *AuthHandler) LoginGin(c *gin.Context) {
	var user model.Login
	validate := validator.New()

	// Decode request body ke struct user
	if err := c.ShouldBindJSON(&user); err != nil {
		l.Log.Error("Invalid request payload", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request payload",
			"error":   err.Error(),
		})
		return
	}

	// Validasi input
	if err := validate.Struct(user); err != nil {
		errors, _ := helper.ValidateInputGeneric(user)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Validation failed",
			"errors":  errors,
		})
		return
	}

	// Proses login
	if err := l.Service.AuthService.Login(&user); err != nil {
		l.Log.Error("Failed to login", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login",
			"error":   err.Error(),
		})
		return
	}

	// Jika sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Login",
		"data":    user,
	})
}

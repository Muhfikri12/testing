package authhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (l *AuthHandler) LoginGin(c *gin.Context) {
	var user model.Login

	if err := c.ShouldBindJSON(&user); err != nil {
		l.Log.Error("Invalid request payload", zap.Error(err))
		helper.ResponsesJson(c, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	if err := l.Service.AuthService.Login(&user); err != nil {
		l.Log.Error("Failed to login", zap.Error(err))
		helper.ResponsesJson(c, http.StatusInternalServerError, "Failed to login", nil)
		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Login", user)
}

package userhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"encoding/json"
	"net/http"

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

func (l *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	user := model.Users{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		l.Log.Error("Invalid request payload: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid request payload: "+err.Error(), nil)
		return
	}

	err := l.Service.UserService.Login(&user)
	if err != nil {
		l.Log.Error("Failed to login: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to login: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Login ", user)
}

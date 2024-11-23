package authhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"ecommers/service"
	"ecommers/util"
	"encoding/json"
	"net/http"

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

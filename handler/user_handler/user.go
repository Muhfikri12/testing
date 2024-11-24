package userhandler

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

func (uh *UserHandler) GetListAddress(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	Addresses, err := uh.Service.UserService.GetListAddress(token)
	if err != nil {
		uh.Log.Error("Data not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data not found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully Get Data", Addresses)
}

func (uh *UserHandler) GetDetailUser(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	users, err := uh.Service.UserService.GetDetailUser(token)
	if err != nil {
		uh.Log.Error("Data not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data not found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully Get Data", users)
}

func (uh *UserHandler) UpdateUserData(w http.ResponseWriter, r *http.Request) {

	User := model.Users{}

	validate := validator.New()
	token := r.Header.Get("Authorization")

	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(User)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(User)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.UpdateUserData(token, &User)
	if err != nil {
		uh.Log.Error("Failed to update Data: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to update Data: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Update Data", User)

}

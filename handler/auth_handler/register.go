package authhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func (u *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := model.Users{}
	validate := validator.New()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.Log.Error("Error from Register Handler :" + err.Error())
		helper.Responses(w, http.StatusBadRequest, "invalid request payload: "+err.Error(), nil)
		return
	}

	err := validate.Struct(user)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(user)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	if err := u.Service.AuthService.Register(&user); err != nil {
		u.Log.Error("Failed to Register: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "failed to Register: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusCreated, "Successfully Register", user)
}

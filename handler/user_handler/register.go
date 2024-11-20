package userhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"encoding/json"
	"net/http"
)

func (u *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	user := model.Users{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.Log.Error("Error from Register Handler :" + err.Error())
		helper.Responses(w, http.StatusBadRequest, "invalid request payload: "+err.Error(), nil)
		return
	}

	if err := u.Service.UserService.Register(&user); err != nil {
		u.Log.Error("Failed to Register: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "failed to Register: "+err.Error(), nil)
		return
	}

	helper.Responses(w, http.StatusCreated, "Successfully Register", user)
}

package userhandler

import (
	"ecommers/helper"
	"net/http"
)

func (uh *UserHandler) GetListAddress(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	users, err := uh.Service.UserService.GetListAddress(token)
	if err != nil {
		uh.Log.Error("Data not found: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data not found", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully Get Data", users)
}

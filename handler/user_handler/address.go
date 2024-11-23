package userhandler

import (
	"ecommers/helper"
	"ecommers/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func (uh *UserHandler) AddAddress(w http.ResponseWriter, r *http.Request) {

	validate := validator.New()
	token := r.Header.Get("Authorization")
	address := model.Addresses{}

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(address)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(address)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.AddAddress(token, &address)
	if err != nil {
		uh.Log.Error("Failed to create address: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to create address", nil)
		return
	}

	helper.Responses(w, http.StatusCreated, "Succesfully Create address", address)
}

func (uh *UserHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {

	validate := validator.New()
	token := r.Header.Get("Authorization")
	address := model.Addresses{}

	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		uh.Log.Error("Invalid Payload Request: " + err.Error())
		helper.Responses(w, http.StatusBadRequest, "Invalid Payload Request", nil)
	}

	err = validate.Struct(address)
	if err != nil {
		errors, _ := helper.ValidateInputGeneric(address)
		helper.Responses(w, http.StatusUnprocessableEntity, "validation failed", errors)
		return
	}

	err = uh.Service.UserService.UpdateAddress(token, &address)
	if err != nil {
		uh.Log.Error("Failed to update address: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to update address", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully Update address", address)
}

func (uh *UserHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	idStr := chi.URLParam(r, "id")

	id, _ := strconv.Atoi(idStr)

	err := uh.Service.UserService.DeleteAddress(token, id)
	if err != nil {
		uh.Log.Error("Failed to delete address: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to update address", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully delete address", nil)
}

func (uh *UserHandler) SetDefault(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("Authorization")

	idStr := chi.URLParam(r, "id")

	id, _ := strconv.Atoi(idStr)

	err := uh.Service.UserService.SetDefault(token, id)
	if err != nil {
		uh.Log.Error("Failed to Set Default address: " + err.Error())
		helper.Responses(w, http.StatusInternalServerError, "Failed to Set Default address", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Succesfully Set Default address", nil)
}

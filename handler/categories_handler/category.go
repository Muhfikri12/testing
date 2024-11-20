package categorieshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"go.uber.org/zap"
)

type CategoriesHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewCategoriesHandler(service service.AllService, log *zap.Logger, config util.Configuration) CategoriesHandler {
	return CategoriesHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (ch *CategoriesHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := ch.Service.CategoryService.GetAllCategories()
	if err != nil {
		ch.Log.Error("Error Handler Product: " + err.Error())
		helper.Responses(w, http.StatusNotFound, "Data tidak tersedia", nil)
		return
	}

	helper.Responses(w, http.StatusOK, "Successfully Get Data", categories)
}

package categorieshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (ch *CategoriesHandler) GetAllCategories(c *gin.Context) {

	categories, err := ch.Service.CategoryService.GetAllCategories()
	if err != nil {
		ch.Log.Error("Error Handler Product: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Data tidak tersedia", nil)

		return
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Get Data", categories)

}

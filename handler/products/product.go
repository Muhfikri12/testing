package productshandler

import (
	"ecommers/helper"
	"ecommers/service"
	"ecommers/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler struct {
	Service service.AllService
	Log     *zap.Logger
	Config  util.Configuration
}

func NewProductsHandler(service service.AllService, log *zap.Logger, config util.Configuration) ProductHandler {
	return ProductHandler{
		Service: service,
		Log:     log,
		Config:  config,
	}
}

func (ph *ProductHandler) GetAll(c *gin.Context) {

	pageStr := c.Query("page")
	category := c.Query("category")
	name := c.Query("name")

	page, _ := strconv.Atoi(pageStr)

	if page < 1 {
		page = 1
	}

	products, totalData, totalPage, err := ph.Service.ProductService.GetAll(page, category, name)
	if err != nil {
		ph.Log.Error("Error Handler Product: " + err.Error())
		helper.ResponsesJson(c, http.StatusNotFound, "Not Found", nil)
		return
	}

	helper.SuccessWithPageGin(c, http.StatusOK, page, 10, totalPage, totalData, "Successfully", products)
}

func (ph *ProductHandler) GetAllBestSelling(c *gin.Context) {

	pageStr := c.Query("page")
	category := c.Query("category")
	name := c.Query("name")

	page, _ := strconv.Atoi(pageStr)

	if page < 1 {
		page = 1
	}

	products, totalData, totalPage, err := ph.Service.ProductService.ProductsBestSelling(page, category, name)
	if err != nil {
		ph.Log.Error("Error Handler Product: " + err.Error())

		helper.ResponsesJson(c, http.StatusNotFound, "Data tidak tersedia", nil)
		return
	}

	helper.SuccessWithPageGin(c, http.StatusOK, page, 10, totalPage, totalData, "Successfully", products)
}

func (ph *ProductHandler) GetProductByID(c *gin.Context) {

	idStr := c.Param("id")

	id, _ := strconv.Atoi(idStr)

	product, err := ph.Service.ProductService.GetProductByID(id)
	if err != nil {
		ph.Log.Error("Error Handler Product: " + err.Error())

		helper.ResponsesJson(c, http.StatusNotFound, "Data tidak tersedia", nil)
	}

	helper.ResponsesJson(c, http.StatusOK, "Successfully Get Data", product)

}

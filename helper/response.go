package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status    int         `json:"status,omitempty"`
	Message   string      `json:"message,omitempty"`
	Page      int         `json:"page,omitempty"`
	Limit     int         `json:"limit,omitempty"`
	TotalPage int         `json:"total_page,omitempty"`
	TotalData int         `json:"total_data,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

func ResponsesJson(c *gin.Context, status int, message string, data any) {

	Response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, Response)
}

func SuccessWithPageGin(c *gin.Context, status, page, limit, totalPage, totalData int, message string, data any) {

	response := Response{
		Status:    status,
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPage,
		Message:   message,
		Data:      data,
	}

	c.JSON(status, response)
}

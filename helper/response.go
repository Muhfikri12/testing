package helper

import (
	"encoding/json"
	"net/http"
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

func Responses(w http.ResponseWriter, code int, message string, data any) {
	response := Response{
		Status:  code,
		Message: message,
		Data:    data,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func SuccessWithPage(w http.ResponseWriter, statusCode, page, limit, totalPage, totalData int, message string, data any) {

	response := Response{
		Status:    statusCode,
		Page:      page,
		Limit:     limit,
		TotalData: totalData,
		TotalPage: totalPage,
		Message:   message,
		Data:      data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

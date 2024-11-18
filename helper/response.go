package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status    int
	Message   string
	Page      int
	Limit     int
	TotalPage int
	TotalData int
	Data      interface{}
}

func Responses(w http.ResponseWriter, code int, message string, data any) {
	response := Response{
		Status:  code,
		Message: message,
		Data:    data,
	}

	w.WriteHeader(code)
	w.Header().Set("content-type", "application/json")
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

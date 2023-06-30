package controller

import (
	"net/http"
	"service/model"

	"github.com/go-chi/render"
)

func badRequest(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set(http.StatusText(http.StatusBadRequest), err.Error())
	res := model.Response{
		Data:    nil,
		Page:    nil,
		Success: false,
		Error:   err.Error(),
	}
	render.JSON(w, r, res)
}

func handlingError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set(http.StatusText(http.StatusInternalServerError), err.Error())
	res := model.Response{
		Data:    nil,
		Page:    nil,
		Success: false,
		Error:   err.Error(),
	}
	render.JSON(w, r, res)
}

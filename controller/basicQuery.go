package controller

import (
	"encoding/json"
	"net/http"
	"service/model"
	"service/service"

	"github.com/go-chi/render"
)

type basicQueryController struct {
	basicQueryService service.BasicQueryService
}

type BasicQueryController interface {
	BasicQuery(w http.ResponseWriter, r *http.Request)
}

// @Summary      Basic Query
// @Description  Basic Query
// @Tags         Basic Query
// @Accept       json
// @Produce      json
// @Param 			 req body model.BasicQueryRequest true "request"
// @Success      200  {object}  model.Response
// @Router       /basic-query/ [post]
func (b *basicQueryController) BasicQuery(w http.ResponseWriter, r *http.Request) {
	var requestData model.BasicQueryRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		badRequest(w, r, err)
		return
	}

	data, err := b.basicQueryService.BasicQuery(requestData)
	if err != nil {
		handlingError(w, r, err)
		return
	}

	res := model.Response{
		Data:    data,
		Success: true,
		Error:   "",
	}

	render.JSON(w, r, res)

}

func NewBasicQueryController() BasicQueryController {
	basicQueryService := service.NewBasicQueryService()
	return &basicQueryController{
		basicQueryService: &basicQueryService,
	}
}

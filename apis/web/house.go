package web

import (
	"encoding/json"
	"github.com/gostarer/domain/infra"
	"github.com/gostarer/domain/infra/base"
	"github.com/gostarer/domain/services"
	"net/http"

	"github.com/gostarer/web"
)

func init() {
	infra.RegisterApi(new(HouseApi))
}

type HouseApi struct {
	service services.HouseService
}

func (h *HouseApi) Init() {
	h.service = services.GetHouseService()
	groupRouter := base.GoStar().Group("/v1/house")
	groupRouter.Post("/create", h.createHandler)
	groupRouter.Get("/get", h.getHandler)
}

func (h *HouseApi) createHandler(ctx *web.Context) {
	var val string
	var err error
	house := services.HouseCreatedDTO{}
	if val = ctx.Param("house"); val != "" {
		ctx.Json(http.StatusInternalServerError, "params error!")
		return
	}
	if err = json.Unmarshal([]byte(val), &house); err != nil {
		ctx.Json(http.StatusInternalServerError, "params json error")
		return
	}
}

func (h *HouseApi) getHandler(ctx *web.Context) {

}

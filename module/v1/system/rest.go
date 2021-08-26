package system

import (
	"go-simple/model"
	"go-simple/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/check", m.check)
}

// @Summary Get Article List
// @Description get all list article
// @Tags Article Management
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/article/list [get]
func (m *Module) check(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    "ok",
		Error:   nil,
	})
}

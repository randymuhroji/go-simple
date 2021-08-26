package auth

import (
	"go-simple/model"
	"go-simple/module/v1/auth/usecase"
	"go-simple/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.POST("/login", m.login)
}

func (m *Module) login(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		payload   = model.LoginPayload{}
	)

	if err := c.Bind(&payload); err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Data:    nil,
			Error:   err,
		})
	}

	resp, err := usecase.Login(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Data:    nil,
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
		Error:   nil,
	})
}

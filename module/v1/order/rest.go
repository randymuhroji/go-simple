package user

import (
	"go-simple/model"
	"go-simple/module/v1/order/usecase"
	"go-simple/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.orderList)
	group.POST("/create", m.orderCreate)
}

func (m *Module) orderList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get order list
	resp, err := usecase.OrderList(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

func (m *Module) orderCreate(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		payload   = model.OrderPayload{}
	)

	err := c.Bind(&payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	payload.CreatedBy = email
	payload.UpdatedBy = email

	resp, err := usecase.CreateOrder(m.Config, &payload)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "order has been created",
		Data:    resp,
	})
}

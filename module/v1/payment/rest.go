package user

import (
	"go-simple/model"
	"go-simple/module/v1/payment/usecase"
	"go-simple/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.paymentList)
	group.POST("/create", m.createPayment)

}

func (m *Module) paymentList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := usecase.PaymentList(m.Config)
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

func (m *Module) createPayment(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		payment   = model.Payment{}
	)

	err := c.Bind(&payment)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	payment.CreatedBy = email
	payment.UpdatedBy = email

	resp, err := usecase.UserRegister(m.Config, &payment)
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
		Message: "user has been registered",
		Data:    payment,
	})
}

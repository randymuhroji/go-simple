package product

import (
	"go-simple/model"
	"go-simple/module/v1/product/usecase"
	"go-simple/utl/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.productList)
	group.GET("/detail/:param", m.productDetail)
	group.PUT("/update/:param", m.productUpdate)
	group.DELETE("/delete/:param", m.userDelete)
	group.POST("/create", m.createProduct)
}

func (m *Module) productList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := usecase.ProductList(m.Config)
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

func (m *Module) createProduct(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		email     = c.Get("email").(string)
		prd       = model.Product{}
	)

	err := c.Bind(&prd)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	prd.CreatedBy = email
	prd.UpdatedBy = email

	resp, err := usecase.CreateProduct(m.Config, &prd)
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
		Data:    resp,
	})
}

func (m *Module) productUpdate(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		prod      = model.Product{}
	)

	id, _ := strconv.Atoi(c.Param("param"))
	prod.Id = id

	err := c.Bind(&prod)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := usecase.UpdateProduct(m.Config, &prod)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been updated",
		Data:    resp,
	})
}

func (m *Module) userDelete(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
		prod      = model.Product{}
	)

	id, _ := strconv.Atoi(c.Param("param"))
	prod.Id = id

	err := c.Bind(&prod)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	_, err = usecase.DeleteProduct(m.Config, &prod)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "your data has been deleted",
		Data:    nil,
	})
}

func (m *Module) productDetail(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)
	id, _ := strconv.Atoi(c.Param("param"))
	resp, err := usecase.ProductDetail(m.Config, id)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
		Error:   err,
	})
}

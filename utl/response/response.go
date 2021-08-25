package response

import (
	"kumparan/model"

	"github.com/labstack/echo/v4"
)

// handling response for success
func Success(c echo.Context, r model.Response) error {
	return c.JSON(r.Status, r)
}

// handling response for error
func Error(c echo.Context, r model.Response) error {
	return c.JSON(r.Status, r)
}

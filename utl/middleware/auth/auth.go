package auth

import (
	"errors"
	"fmt"
	"go-simple/config"
	"go-simple/model"
	glJwt "go-simple/utl/jwt"
	"go-simple/utl/response"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Handle struct {
	Config config.Configuration
}

func InitAuthMiddleware(conf config.Configuration) *Handle {
	return &Handle{
		Config: conf,
	}
}

// BearerClaims data structure for claims
type BearerClaims struct {
	Email    string    `json:"email"`
	Hostname string    `json:"hostname"`
	Exp      int64     `json:"exp"`
	Iat      time.Time `json:"iat"`
	jwt.StandardClaims
}

func (h *Handle) BearerVerify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")

			// validation authorization
			if len(auth) <= 0 {
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: nil,
					Data:    nil,
					Error:   "authorization is empty",
				})
			}

			splitToken := strings.Split(auth, " ")
			if len(splitToken) < 2 {
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: nil,
					Data:    nil,
					Error:   "authorization is empty",
				})
			}

			if splitToken[0] != "Bearer" {
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: nil,
					Data:    nil,
					Error:   "authorization is empty",
				})
			}

			tokenStr := splitToken[1]

			tkn, err := jwt.ParseWithClaims(tokenStr, &glJwt.BearerClaims{}, func(t *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != t.Method {
					return nil, errors.New("system unauthorized")
				}

				return []byte(glJwt.Key), nil
			})

			if err != nil {
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: err.Error(),
					Data:    nil,
					Error:   "invalid authorization",
				})
			}

			if claims, ok := tkn.Claims.(*glJwt.BearerClaims); tkn.Valid && ok {

				if !time.Unix(int64(claims.Exp), 0).After(time.Now()) || time.Unix(int64(claims.Exp), 0).Before(time.Now()) {
					return response.Error(c, model.Response{
						LogId:   c.Get("request_id").(string),
						Status:  http.StatusUnauthorized,
						Message: nil,
						Data:    nil,
						Error:   "token has been expired",
					})
				}
				c.Set("token", tkn.Raw)
				c.Set("email", claims.Email)

				return next(c)

			} else if ve, ok := err.(*jwt.ValidationError); ok {

				var errorStr string
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					errorStr = "Token has been expired"
				} else {
					errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
				}
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: errorStr,
					Data:    nil,
					Error:   "invalid authorization",
				})
			} else {
				return response.Error(c, model.Response{
					LogId:   c.Get("request_id").(string),
					Status:  http.StatusUnauthorized,
					Message: err.Error(),
					Data:    nil,
					Error:   "invalid authorization",
				})
			}

			// return next(c)

		}
	}
}

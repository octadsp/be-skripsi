package middleware

import (
	dto "be-skripsi/dto/results"
	jwtToken "be-skripsi/pkg/jwt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Auth Function
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, dto.ErrorResult{Status: http.StatusBadRequest, Message: "unauthorized"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, Result{Status: http.StatusUnauthorized, Message: "unauthorized"})
		}

		c.Set("userLogin", claims)
		return next(c)
	}
}

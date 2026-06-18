package middleware

import (
	"net/http"

	"dinz-rentbike/pkg/response"

	echo "github.com/labstack/echo/v4"
)

func RoleMiddleware(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("user_role")
			if userRole == nil || userRole.(string) != role {
				return response.ErrorResponse(c, http.StatusForbidden, "akses tidak diizinkan")
			}

			return next(c)
		}
	}
}

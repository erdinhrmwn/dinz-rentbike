package middleware

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"dinz-rentbike/pkg/response"
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

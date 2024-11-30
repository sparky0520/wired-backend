package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase" // Core PocketBase package
	// Provides models like Record for working with database records
	// For handling API-related functionality
)

func AuthMiddleware(app *pocketbase.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extracting authorization header
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			// extract token
			token := strings.TrimPrefix(authHeader, "Bearer ")

			// verify token with Pocketbase
			record, err := app.FindAuthRecordByToken(token)
			if err != nil || record == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			}

			// user is authenticated - store user record for later use
			c.Set("user", record)

			// continue to next handler
			return next(c)
		}
	}
}

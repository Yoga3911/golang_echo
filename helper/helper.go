package helper

import "github.com/labstack/echo"

func Helper(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"messages": message,
		"user":     data,
	})
}

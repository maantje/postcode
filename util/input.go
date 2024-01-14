package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParamInt64(c echo.Context, name string, def int64) int64 {
	param := c.Param(name)
	result, err := strconv.Atoi(param)
	if err != nil {
		return def
	}

	return int64(result)
}

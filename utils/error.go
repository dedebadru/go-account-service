package utils

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func RespondWithError(c echo.Context, statusCode int, message string) error {
    return c.JSON(statusCode, echo.Map{
        "remark": message,
    })
}

func BadRequestError(c echo.Context, message string) error {
    return RespondWithError(c, http.StatusBadRequest, message)
}

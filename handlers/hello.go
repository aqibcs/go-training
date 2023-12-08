package handlers

import (
	"github.com/labstack/echo/v4"
	"go-training/models/request"
	"go-training/models/response"
	"net/http"
	"time"
)

func HelloHandler(c echo.Context) error {
	var requestBody request.RequestBody
	err := c.Bind(&requestBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	responseBody := response.ResponseBody{
		Code:      200,
		Message:   "Welcome " + requestBody.Name + "!",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, responseBody)
}

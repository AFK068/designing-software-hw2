package hsezooapi

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/labstack/echo/v4"

	hsezootypes "github.com/AFK068/hsezoo/internal/api/openapi/hsezoo/v1"
)

func SendSuccessResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, data)
}

func SendBadRequestResponse(ctx echo.Context, err, description string) error {
	return ctx.JSON(http.StatusBadRequest, hsezootypes.ApiErrorResponse{
		Description:      aws.String(description),
		Code:             aws.String("400"),
		ExceptionMessage: aws.String(err),
	})
}

func SendNotFoundResponse(ctx echo.Context, err, description string) error {
	return ctx.JSON(http.StatusNotFound, hsezootypes.ApiErrorResponse{
		Description:      aws.String(description),
		Code:             aws.String("404"),
		ExceptionMessage: aws.String(err),
	})
}

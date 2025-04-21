package responses

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code      int       `json:"code"`
	Success   bool      `json:"success"`
	Data      any       `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

func NewApplicationResponse[T any](ctx echo.Context, code int, data T, isSuccess bool) error {
	return ctx.JSON(code, Response{
		Code:      code,
		Success:   isSuccess,
		Data:      data,
		Timestamp: time.Now(),
	})
}

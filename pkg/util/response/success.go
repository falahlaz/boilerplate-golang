package response

import (
	"net/http"

	"github.com/falahlaz/boilerplate-golang/internal/abstraction"
	"github.com/labstack/echo/v4"
)

type successResponse struct {
	Meta `json:"meta"`
	Data any `json:"data"`
}

type Success struct {
	Response successResponse
	Code     int
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}

type successConstant struct {
	OK Success
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Meta: Meta{
				Message: "Request successfully proceed",
				Info:    &abstraction.PaginationInfo{},
			},
			Data: nil,
		},
		Code: http.StatusOK,
	},
}

func SuccessBuilder(res *Success, data any) *Success {
	res.Response.Data = data
	return res
}

func SuccessResponse(data any) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func CustomSuccessBuilder(code int, message string, info *abstraction.PaginationInfo, data any) *Success {
	return &Success{
		Response: successResponse{
			Meta: Meta{
				Message: message,
				Info:    info,
			},
			Data: data,
		},
		Code: code,
	}
}

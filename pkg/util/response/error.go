package response

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Meta  `json:"meta"`
	Error string
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

func (e *Error) Error() string {
	return e.ErrorMessage.Error()
}

func (e *Error) Send(c echo.Context) error {
	if e.ErrorMessage != nil {
		logrus.Error(e.ErrorMessage)
	}
	return c.JSON(e.Code, e.Response)
}

const (
	E_DUPLICATE_ENTITY     = "duplicate_entity"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	NotFound                 Error
	RouteNotFound            Error
	UnprocessableEntity      Error
	DuplicateEntity          Error
	Unauthorized             Error
	BadRequest               Error
	Validation               Error
	InternalServerError      Error
	EmailOrPasswordIncorrect Error
}

var ErrorConstant errorConstant = errorConstant{
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	DuplicateEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Data already exist",
			},
			Error: E_DUPLICATE_ENTITY,
		},
		Code: http.StatusConflict,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Unauthorized",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Bad request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Internal server error",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	EmailOrPasswordIncorrect: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Email or password is incorrect",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
}

func ErrorBuilder(err *Error, message error, customMessage ...string) *Error {
	err.ErrorMessage = message

	if strings.Contains(strings.Join([]string{E_UNPROCESSABLE_ENTITY, E_BAD_REQUEST, E_DUPLICATE_ENTITY}, ","), err.Response.Error) {
		err.Response.Meta.Message = message.Error()
	}
	if len(customMessage) > 0 {
		strings.Join(customMessage, " \n")
	}

	return err
}

func CustomErrorBuilder(code int, message string, err error) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err.Error(),
		},
		Code:         code,
		ErrorMessage: err,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	}
	return ErrorBuilder(&ErrorConstant.InternalServerError, err)
}

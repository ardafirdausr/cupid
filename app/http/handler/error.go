package handler

import (
	"com.ardafirdausr.cupid/app/http/handler/response"
	"com.ardafirdausr.cupid/internal/entity/errs"
	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	message := echo.ErrInternalServerError.Message.(string)
	response := response.BasicResponse{
		Message: message,
		Data:    nil,
	}

	if he, ok := err.(*echo.HTTPError); ok {
		response.Message = he.Message.(string)
		ctx.JSON(he.Code, response)
		return
	}

	response.Message = errs.GetCauserMessage(err, message)
	ctx.JSON(errs.GetHttpCode(err), response)
}

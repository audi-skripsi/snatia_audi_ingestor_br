package httputil

import (
	"net/http"

	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/dto"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/errors"
	"github.com/labstack/echo/v4"
)

func WriteSuccessResponse(w echo.Context, payload interface{}) error {
	return WriteResponse(w, dto.ResponseParam{
		Status: http.StatusOK,
		Payload: dto.BaseResponse{
			Data: payload,
		},
	})
}

func WriteErrorResponse(w echo.Context, er error) error {
	errResp := errors.GetErrorResponse(er)
	return WriteResponse(w, dto.ResponseParam{
		Status: int(errResp.Code),
		Payload: dto.BaseResponse{
			Error: &dto.ErrorResponse{
				Code:    errResp.Code,
				Message: errResp.Message,
			},
		},
	})
}

func WriteResponse(w echo.Context, param dto.ResponseParam) error {
	return w.JSON(param.Status, param.Payload)
}

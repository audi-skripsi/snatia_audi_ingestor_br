package handler

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/dto"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/errors"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/util/httputil"
	"github.com/labstack/echo/v4"
)

type EventHandlerService = func(event dto.EventLog, microBatch bool) (err error)

func HandleNonMicrobatch(service EventHandlerService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request dto.EventLog

		if err = c.Bind(&request); err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrBadRequest)
		}

		err = service(request, false)

		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, "ok")
	}
}

func HandleMicrobatch(service EventHandlerService) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var request dto.EventLog

		if err = c.Bind(&request); err != nil {
			return httputil.WriteErrorResponse(c, errors.ErrBadRequest)
		}

		err = service(request, true)

		if err != nil {
			return httputil.WriteErrorResponse(c, err)
		}

		return httputil.WriteSuccessResponse(c, "ok")
	}
}

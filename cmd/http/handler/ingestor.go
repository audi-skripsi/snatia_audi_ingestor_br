package handler

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/pkg/util/httputil"
	"github.com/labstack/echo/v4"
)

func HandleNonMicrobatch() echo.HandlerFunc {
	return func(c echo.Context) error {

		return httputil.WriteSuccessResponse(c, "ok")
	}
}

func HandleMicrobatch() echo.HandlerFunc {
	return func(c echo.Context) error {

		return httputil.WriteSuccessResponse(c, "ok")
	}
}

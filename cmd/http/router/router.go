package router

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/cmd/http/handler"
	"github.com/labstack/echo/v4"
)

type RouterInitParams struct {
	Ec *echo.Echo
}

func Init(params *RouterInitParams) {
	params.Ec.POST(MicrobatchPath, handler.HandleMicrobatch())
	params.Ec.POST(MicrobatchPath, handler.HandleNonMicrobatch())
}

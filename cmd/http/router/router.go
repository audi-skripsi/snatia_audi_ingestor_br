package router

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/cmd/http/handler"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/service"
	"github.com/labstack/echo/v4"
)

type RouterInitParams struct {
	Ec      *echo.Echo
	Service service.Service
}

func Init(params *RouterInitParams) {
	params.Ec.POST(MicrobatchPath, handler.HandleMicrobatch(params.Service.StoreEvent))
	params.Ec.POST(NonMicrobatchPath, handler.HandleNonMicrobatch(params.Service.StoreEvent))
}

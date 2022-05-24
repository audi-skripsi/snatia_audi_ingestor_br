package http

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/cmd/http/router"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ServerInitParams struct {
	Logger  *logrus.Entry
	Config  *config.Config
	Ec      *echo.Echo
	Service service.Service
}

func StartServer(param *ServerInitParams) {
	router.Init(&router.RouterInitParams{
		Ec:      param.Ec,
		Service: param.Service,
	})

	err := param.Ec.Start(param.Config.AppAddress)
	if err != nil {
		param.Logger.Errorf("error starting server at %s: %+v", param.Config.AppAddress, err)
	}
}

package http

import (
	"github.com/audi-skripsi/snatia_audi_ingestor_be/cmd/http/router"
	"github.com/audi-skripsi/snatia_audi_ingestor_be/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ServerInitParams struct {
	Logger *logrus.Entry
	Config *config.Config
	Ec     *echo.Echo
}

func StartServer(param *ServerInitParams) {
	router.Init(&router.RouterInitParams{
		Ec: param.Ec,
	})

	err := param.Ec.Start(param.Config.AppAddress)
	if err != nil {
		param.Logger.Errorf("error starting server at %s: %+v", param.Config.AppAddress, err)
	}
}

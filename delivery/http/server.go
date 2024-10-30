package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"mlm/config"
	"mlm/delivery/http/handler/node_handler"
	_ "mlm/docs"
	"mlm/logger"
	"mlm/service/node_svc"
	"mlm/service/user_svc"
	"mlm/validator/node_validator"
)

type Server struct {
	config      config.Application
	nodeHandler node_handler.NodeHandler
	Router      *echo.Echo
}

func New(
	config config.Application,
	nodeSvc node_svc.NodeService,
	userSvc user_svc.Service,
	nodeValidator node_validator.Validator,
) Server {
	return Server{
		config:      config,
		nodeHandler: node_handler.NewNodeHandler(nodeSvc, nodeValidator, userSvc),
		Router:      echo.New(),
	}
}

func (s Server) Serve() {

	s.Router.Use(middleware.RequestID())

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""

			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.Logger.Info("test", zap.Any("all", v))

			logger.Logger.Named("http-server").Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))

	s.Router.Use(middleware.Recover())

	s.nodeHandler.SetRouter(s.Router)
	s.Router.GET("/swagger/*", echoSwagger.WrapHandler)

	address := fmt.Sprintf(":%s", s.config.HTTPServer.Port)
	logger.Logger.Info("echo", zap.Any("start echo server on", address))

	if err := s.Router.Start(address); err != nil {
		logger.Logger.Info("echo", zap.Any("server error", err))
	}

}

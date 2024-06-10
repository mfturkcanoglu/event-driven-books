package server

import (
	"book_publisher/internal/config"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	logger       *zap.Logger
	config       *config.Config
	echoInstance *echo.Echo
}

func NewServer(logger *zap.Logger, config *config.Config) *Server {
	return &Server{
		logger: logger,
		config: config,
	}
}

func (server *Server) Serve() {
	echoInstance := echo.New()
	server.echoInstance = echoInstance

	server.addMiddlewaresToInstance()

	serveError := echoInstance.Start(server.config.Port)
	server.logger.Fatal("Server is down", zap.Error(serveError))
}

func (server *Server) addMiddlewaresToInstance() {
	server.echoInstance.Use(server.getZapLoggerMiddleware())
}

func (server *Server) getZapLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogError:   true,
		LogMethod:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			server.logger.Info("request",
				zap.Duration("latency", v.Latency),
				zap.Error(v.Error),
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.String("date", v.StartTime.Format(time.RFC3339)),
			)

			return nil
		},
	})
}

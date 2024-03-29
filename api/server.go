package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/samthehai/ml-backend-test-samthehai/config"
	"github.com/samthehai/ml-backend-test-samthehai/pkg/logger"
)

const ctxTimeout = 5

type ConnManager interface {
	GetReader() *sqlx.DB
	GetWriter() *sqlx.DB
	CloseAll()
}

type Server struct {
	echo        *echo.Echo
	cfg         *config.Config
	logger      logger.Logger
	connManager ConnManager
}

func NewServer(cfg *config.Config, logger logger.Logger, connManager ConnManager) *Server {
	return &Server{
		echo:        echo.New(),
		cfg:         cfg,
		logger:      logger,
		connManager: connManager,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr: s.cfg.Server.Port,
	}

	go func() {
		s.logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting Server: ", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}

package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type server struct {
	httpServer *http.Server
}

func New(port string, handler http.Handler) *server {
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return &server{httpServer: httpServer}
}

func (s *server) Run(ctx context.Context) error {
	var grp errgroup.Group
	grp.Go(func() error {
		if err := s.listenAndServe(); err != nil {
			return errors.New("server error listening")
		}
		return nil
	})

	grp.Go(func() error {
		if err := s.waitAndShutdown(ctx); err != nil && err != http.ErrServerClosed {
			return errors.New("server shutdown error")
		}
		return nil
	})

	return grp.Wait()
}

func (s *server) listenAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) waitAndShutdown(ctx context.Context) error {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
	return s.httpServer.Shutdown(ctx)
}

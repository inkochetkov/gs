package gs

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/inkochetkov/log"
)

func TestGS(t *testing.T) {

	logger := log.New(log.DevLog, "", "")

	mainCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	module := New(mainCtx, logger)

	srv := new(logger)

	module.Add(srv.Start(module.ctx), srv.Stop(module.ctx))

	module.Start()

}

type serv struct {
	srv    *http.Server
	logger *log.Log
}

func new(logger *log.Log) *serv {
	return &serv{
		srv: &http.Server{
			Addr: ":8800",
		},
		logger: logger,
	}
}

func (s *serv) Start(ctx context.Context) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		if err := s.srv.ListenAndServe(); err != nil {
			select {
			case <-ctx.Done():
				return nil
			default:
				s.logger.Error("http ", err)
				return err
			}
		}
		return nil
	}
}

func (s *serv) Stop(ctx context.Context) func(ctx context.Context) error {
	return func(ctx context.Context) error {

		if err := s.srv.Close(); err != nil {
			s.logger.Error(Stop, err)
			return err
		}

		return nil
	}
}

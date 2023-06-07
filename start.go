package gs

import (
	"context"

	"github.com/inkochetkov/log"

	"golang.org/x/sync/errgroup"
)

const (
	App   = "app: "
	Start = "start"
	Stop  = "stop"
)

type Module struct {
	logger *log.Log
	ctx    context.Context
	eg     *errgroup.Group
}

func New(ctx context.Context, logger *log.Log) *Module {

	errGroup, moduleCtx := errgroup.WithContext(NewModuleContext(ctx))

	return &Module{
		ctx:    moduleCtx,
		logger: logger,
		eg:     errGroup,
	}
}

// Start all module
func (s *Module) Start() {

	s.logger.Info(App, Start)
	defer s.logger.Info(App, Stop)

	if err := s.eg.Wait(); err != nil {
		s.logger.Error(App, err)
	}

}

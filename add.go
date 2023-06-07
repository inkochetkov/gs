package gs

import "context"

// Add module
func (s *Module) Add(fStart func(ctx context.Context) error, fStop func(ctx context.Context) error) {

	// start module
	s.eg.Go(func() error {

		return fStart(s.ctx)

	})

	// stop module
	s.eg.Go(func() error {

		<-s.ctx.Done()

		return fStop(s.ctx)
	})
}

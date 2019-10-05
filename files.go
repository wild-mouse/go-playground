package filesapi

import (
	"context"
	"log"

	files "github.com/wild-mouse/go-playground/gen/files"
)

// files service example implementation.
// The example methods log the requests and return zero values.
type filessrvc struct {
	logger *log.Logger
}

// NewFiles returns the files service implementation.
func NewFiles(logger *log.Logger) files.Service {
	return &filessrvc{logger}
}

// Add implements add.
func (s *filessrvc) Add(ctx context.Context, p *files.AddPayload) (err error) {
	s.logger.Print("files.add")
	s.logger.Print(*p.File)
	return
}

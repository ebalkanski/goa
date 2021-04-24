package service

import (
	"context"
	"log"

	"github.com/ebalkanski/goa/gen/play"
)

type playsvc struct {
	logger *log.Logger
}

// NewPlay returns the play service implementation.
func NewPlay(logger *log.Logger) play.Service {
	return &playsvc{logger}
}

// Add implements add
func (s *playsvc) Add(ctx context.Context, p *play.AddPayload) (res int, err error) {
	s.logger.Print("play.add")

	return p.A + p.B, nil
}

// Rate implements rate
func (s *playsvc) Rate(ctx context.Context, payload *play.RatePayload) (err error) {
	s.logger.Printf("%#v\n", payload)
	s.logger.Printf("%#v\n", *payload.ID)
	s.logger.Printf("%#v\n", payload.Rates)

	return nil
}

package covid19info

import (
	"context"
	"covid-19-alert-to-slack/internal/entity"
	"log"
)

type covid19Info struct {
	entity.Covid19InfoEntity
}

type service struct {
	repo   Repository
	logger log.Logger
}

// Service encapsulates usecase logic for covid19Info.
type Service interface {
	Get(ctx context.Context, id string) (covid19Info, error)
}

// NewService creates a covid19info service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(ctx context.Context, id string) (covid19Info, error) {
	panic("implement me")
}

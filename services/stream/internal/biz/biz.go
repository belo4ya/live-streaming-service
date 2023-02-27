package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewStreamUseCase)

type Stream struct {
	Id       int
	Name     string
	Username string
}

type StreamRepo interface {
	ListAll(context.Context) ([]*Stream, error)
}

type StreamUseCase struct {
	repo StreamRepo
	log  *log.Helper
}

func NewStreamUseCase(repo StreamRepo, logger log.Logger) *StreamUseCase {
	return &StreamUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *StreamUseCase) ListAll(ctx context.Context) ([]*Stream, error) {
	return uc.repo.ListAll(ctx)
}

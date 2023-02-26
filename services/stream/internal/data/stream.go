package data

import (
	"context"
	"github.com/belo4ya/live-streaming-service/services/stream/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type streamRepo struct {
	data *Data
	log  *log.Helper
}

func NewStreamRepo(data *Data, logger log.Logger) biz.StreamRepo {
	return &streamRepo{data: data, log: log.NewHelper(logger)}
}

func (r streamRepo) ListAll(context.Context) ([]*biz.Stream, error) {
	return nil, nil
}

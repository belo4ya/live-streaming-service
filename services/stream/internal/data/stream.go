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

func (r *streamRepo) ListAll(context.Context) ([]*biz.Stream, error) {
	streams := []*biz.Stream{
		{
			Id:       1,
			Name:     "Stream 1",
			Username: "Streamer 1",
		},
		{
			Id:       2,
			Name:     "Stream 2",
			Username: "Streamer 2",
		},
		{
			Id:       3,
			Name:     "Stream 3",
			Username: "Streamer 3",
		},
	}
	return streams, nil
}

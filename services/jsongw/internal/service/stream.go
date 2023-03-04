package service

import (
	"context"
	"encoding/json"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
)

type StreamService struct {
	v1.UnimplementedStreamServiceServer
	data *data.Data
	log  *log.Helper
}

func NewStreamService(data *data.Data, logger log.Logger) *StreamService {
	return &StreamService{data: data, log: log.NewHelper(logger)}
}

func (s *StreamService) ListStreams(ctx context.Context, req *empty.Empty) (*v1.ListStreamsResponse, error) {
	resp, err := s.data.StreamC.ListStreams(ctx, req)
	if err != nil {
		return nil, err
	}
	return proxy(&v1.ListStreamsResponse{}, resp)
}

func proxy[D any, S any](dst D, src S) (D, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return *new(D), err
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		return *new(D), err
	}
	return dst, nil
}

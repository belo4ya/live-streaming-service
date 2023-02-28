package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	"github.com/belo4ya/live-streaming-service/services/jsongw/internal/data"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/proto"
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
	r, err := s.data.StreamC.ListStreams(ctx, req)
	if err != nil {
		return nil, err
	}
	return proto.Clone(r).(*v1.ListStreamsResponse), nil
}

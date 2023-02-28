package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	streamv1 "github.com/belo4ya/live-streaming-service/api/stream/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type StreamService struct {
	v1.UnimplementedStreamServiceServer
	c streamv1.StreamServiceClient
}

func NewStreamService(c streamv1.StreamServiceClient) *StreamService {
	return &StreamService{c: c}
}

func (s *StreamService) ListStreams(ctx context.Context, req *empty.Empty) (*v1.ListStreamsResponse, error) {
	r, err := s.c.ListStreams(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListStreamsResponse{
		Results: make([]*v1.ListStreamsResponse_Stream, 0),
	}
	for _, r := range r.Results {
		res := &v1.ListStreamsResponse_Stream{
			Id:       r.Id,
			Name:     r.Name,
			Username: r.Username,
		}
		resp.Results = append(resp.Results, res)
	}
	return &v1.ListStreamsResponse{Results: resp.Results}, nil
}

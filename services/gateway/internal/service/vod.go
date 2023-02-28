package service

import (
	"context"
	v1 "github.com/belo4ya/live-streaming-service/api/jsongw/v1"
	vodv1 "github.com/belo4ya/live-streaming-service/api/vod/v1"
	"github.com/golang/protobuf/ptypes/empty"
)

type VODService struct {
	v1.UnimplementedVODServiceServer
	c vodv1.VODServiceClient
}

func NewVODService(c vodv1.VODServiceClient) *VODService {
	return &VODService{c: c}
}

func (s *VODService) ListVODs(ctx context.Context, req *empty.Empty) (*v1.ListVODsResponse, error) {
	r, err := s.c.ListVODs(ctx, req)
	if err != nil {
		return nil, err
	}

	resp := &v1.ListVODsResponse{
		Results: make([]*v1.ListVODsResponse_VOD, 0),
	}
	for _, r := range r.Results {
		res := &v1.ListVODsResponse_VOD{
			Id:        r.Id,
			Name:      r.Name,
			Username:  r.Username,
			CreatedAt: r.CreatedAt,
		}
		resp.Results = append(resp.Results, res)
	}
	return resp, nil
}

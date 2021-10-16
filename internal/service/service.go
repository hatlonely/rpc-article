package service

import (
	"context"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
)

type Options struct {
}

type Service struct {
	options *Options

	api.UnimplementedArticleServiceServer
}

func NewServiceWithOptions(options *Options) (*Service, error) {
	return &Service{
		options: options,
	}, nil
}

func (s *Service) Ping(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

package service

import (
	"context"

	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/storage"
)

type Options struct {
	Storage refx.TypeOptions
}

type Service struct {
	options *Options
	storage storage.Storage

	api.UnimplementedArticleServiceServer
}

func NewServiceWithOptions(options *Options, opts ...refx.Option) (*Service, error) {
	s, err := storage.NewStorageWithOptions(&options.Storage, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "storage.NewStorageWithOptions failed")
	}

	return &Service{
		options: options,
		storage: s,
	}, nil
}

func (s *Service) Ping(ctx context.Context, req *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}

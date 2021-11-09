package service

import (
	"context"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
)

func (s *Service) DelAuthorByKey(ctx context.Context, req *api.Author) (*api.Empty, error) {
	if err := s.storage.DelAuthorByKey(ctx, req.Key); err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

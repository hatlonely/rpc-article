package service

import (
	"context"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
)

func (s *Service) DelAuthor(ctx context.Context, req *api.AuthorID) (*api.Empty, error) {
	if err := s.storage.DelAuthor(ctx, req.Id); err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

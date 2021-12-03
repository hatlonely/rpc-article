package service

import (
	"context"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/storage"
)

func (s *Service) PutAuthor(ctx context.Context, req *api.Author) (*api.AuthorID, error) {
	id, err := s.storage.PutAuthor(ctx, &storage.Author{
		Key:    req.Key,
		Name:   req.Name,
		Avatar: req.Avatar,
	})
	if err != nil {
		return nil, err
	}

	return &api.AuthorID{Id: id}, nil
}

func (s *Service) DelAuthor(ctx context.Context, req *api.AuthorID) (*api.Empty, error) {
	if err := s.storage.DelAuthor(ctx, req.Id); err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

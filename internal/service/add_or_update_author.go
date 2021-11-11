package service

import (
	"context"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/storage"
)

func (s *Service) AddOrUpdateAuthor(ctx context.Context, req *api.Author) (*api.AuthorID, error) {
	author, err := s.storage.GetAuthorByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	if author == nil {
		id, err := s.storage.PutAuthor(ctx, &storage.Author{
			Key:  req.Key,
			Name: req.Name,
		})
		if err != nil {
			return nil, err
		}
		return &api.AuthorID{
			Id: id,
		}, nil
	}

	if err := s.storage.UpdateAuthor(ctx, &storage.Author{
		ID:   author.ID,
		Key:  req.Key,
		Name: req.Name,
	}); err != nil {
		return nil, err
	}

	return &api.AuthorID{
		Id: author.ID,
	}, nil
}

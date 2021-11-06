package service

import (
	"context"
	"strings"

	"github.com/hatlonely/go-kit/rpcx"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/storage"
)

func (s *Service) PutArticle(ctx context.Context, req *api.Article) (*api.ArticleID, error) {
	id, err := s.storage.PutArticle(ctx, &storage.Article{
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Tags:     strings.Join(req.Tags, "|"),
		Brief:    req.Brief,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &api.ArticleID{Id: id}, nil
}

func (s *Service) GetArticle(ctx context.Context, req *api.ArticleID) (*api.Article, error) {
	article, err := s.storage.GetArticle(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, rpcx.NewErrorf(nil, codes.NotFound, "ArticleNotExists", "article not exist")
	}
	return &api.Article{
		Id:       article.ID,
		AuthorID: article.AuthorID,
		Title:    article.Title,
		Tags:     strings.Split(article.Tags, "|"),
		Brief:    article.Brief,
		Content:  article.Content,
		CreateAt: int32(article.CreateAt.Unix()),
		UpdateAt: int32(article.UpdateAt.Unix()),
	}, nil
}

func (s *Service) DelArticle(ctx context.Context, req *api.ArticleID) (*api.Empty, error) {
	err := s.storage.DelArticle(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *Service) UpdateArticle(ctx context.Context, req *api.Article) (*api.Empty, error) {
	err := s.storage.UpdateArticle(ctx, &storage.Article{
		ID:       req.Id,
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Tags:     strings.Join(req.Tags, "|"),
		Brief:    req.Brief,
		Content:  req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *Service) ListArticle(ctx context.Context, req *api.ListArticleReq) (*api.ListArticleRes, error) {
	articles, err := s.storage.ListArticles(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	if articles == nil {
		return nil, rpcx.NewErrorf(nil, codes.NotFound, "ArticleNotExists", "article not exist")
	}

	var res api.ListArticleRes
	for _, article := range articles {
		res.Articles = append(res.Articles, &api.Article{
			Id:       article.ID,
			AuthorID: article.AuthorID,
			Title:    article.Title,
			Tags:     strings.Split(article.Tags, "|"),
			Brief:    article.Brief,
			Content:  article.Content,
			CreateAt: int32(article.CreateAt.Unix()),
			UpdateAt: int32(article.UpdateAt.Unix()),
		})
	}

	return &res, nil
}

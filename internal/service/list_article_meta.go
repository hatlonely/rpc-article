package service

import (
	"context"
	"strings"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/rpc-article/api/gen/go/api"
	"github.com/hatlonely/rpc-article/internal/storage"
	"google.golang.org/grpc/codes"
)

func (s *Service) ListArticleMeta(ctx context.Context, req *api.ListArticleMetaReq) (*api.ListArticleMetaRes, error) {
	if req.Limit == 0 {
		req.Limit = 100
	}

	var articles []*storage.Article
	var err error
	if len(req.AuthorID) == 0 {
		articles, err = s.storage.ListArticles(ctx, req.Offset, req.Limit)
	} else {
		articles, err = s.storage.ListArticlesByAuthor(ctx, req.AuthorID, req.Offset, req.Limit)
	}
	if err != nil {
		return nil, err
	}
	if articles == nil {
		return nil, rpcx.NewErrorf(nil, codes.NotFound, "ArticleNotExists", "article not exist")
	}

	var res api.ListArticleMetaRes
	for _, article := range articles {
		res.ArticleMetas = append(res.ArticleMetas, &api.ArticleMeta{
			Id:       article.ID,
			AuthorID: article.AuthorID,
			Title:    article.Title,
			Tags:     strings.Split(article.Tags, "|"),
			Brief:    article.Brief,
			CreateAt: int32(article.CreateAt.Unix()),
			UpdateAt: int32(article.UpdateAt.Unix()),
		})
	}

	return &res, nil
}

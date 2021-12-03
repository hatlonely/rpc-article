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

	articleJoinAuthors, err := s.storage.ListArticleJoinAuthor(ctx, req.AuthorID, req.Offset, req.Limit)
	if err != nil {
		if err == storage.ErrNotFound {
			return nil, rpcx.NewErrorf(nil, codes.NotFound, "ArticleNotExists", "article not exist")
		}
		return nil, err
	}

	var res api.ListArticleMetaRes
	for _, articleJoinAuthor := range articleJoinAuthors {
		res.ArticleMetas = append(res.ArticleMetas, &api.ArticleMeta{
			Id:         articleJoinAuthor.Article.ID,
			AuthorID:   articleJoinAuthor.Article.AuthorID,
			Title:      articleJoinAuthor.Article.Title,
			Tags:       strings.Split(articleJoinAuthor.Article.Tags, "|"),
			Brief:      articleJoinAuthor.Article.Brief,
			CreateAt:   int32(articleJoinAuthor.Article.CreateAt.Unix()),
			UpdateAt:   int32(articleJoinAuthor.Article.UpdateAt.Unix()),
			AuthorName: articleJoinAuthor.Author.Name,
		})
	}

	return &res, nil
}

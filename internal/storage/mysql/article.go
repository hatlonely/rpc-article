package mysql

import (
	"context"
	"encoding/hex"

	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-article/internal/storage"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (m *MySQL) PutArticle(ctx context.Context, article *storage.Article) (string, error) {
	article.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	if err := m.db.Clauses(ctx, clause.OnConflict{
		UpdateAll: true,
	}).Create(ctx, article).Unwrap().Error; err != nil {
		return "", errors.Wrap(err, "db.Create failed")
	}
	var newArticle storage.Article
	if err := m.db.Where(ctx, &storage.Article{
		AuthorID: article.AuthorID,
		Title:    article.Title,
	}).First(ctx, &newArticle).Unwrap().Error; err != nil {
		return "", errors.Wrap(err, "db.Find failed")
	}
	return newArticle.ID, nil
}

func (m *MySQL) GetArticle(ctx context.Context, id string) (*storage.Article, error) {
	var article storage.Article
	if err := m.db.
		Where(ctx, "`id`=?", id).
		First(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, storage.ErrNotFound
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQL) GetArticleByAuthorAndTitle(ctx context.Context, authorID string, title string) (*storage.Article, error) {
	var article storage.Article
	if err := m.db.
		Where(ctx, &storage.Article{AuthorID: authorID, Title: title}).
		Find(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, storage.ErrNotFound
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQL) UpdateArticle(ctx context.Context, article *storage.Article) error {
	return m.db.Model(ctx, article).Where(ctx, "`id`=?", article.ID).Updates(ctx, article).Unwrap().Error
}

func (m *MySQL) DelArticle(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Article{ID: id}).Unwrap().Error
}

func (m *MySQL) ListArticles(ctx context.Context, offset int32, limit int32) ([]*storage.Article, error) {
	var articles []*storage.Article

	if err := m.db.
		Order(ctx, "`createAt` DESC, `authorID`, `title`").
		Offset(ctx, int(offset)).
		Limit(ctx, int(limit)).
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

func (m *MySQL) ListArticlesByAuthor(ctx context.Context, authorID string, offset int32, limit int32) ([]*storage.Article, error) {
	var articles []*storage.Article

	if err := m.db.
		Where(ctx, "`authorID`=?", authorID).
		Order(ctx, "`createAt` DESC, `authorID`, `title`").
		Offset(ctx, int(offset)).
		Limit(ctx, int(limit)).
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

func (m *MySQL) ListArticlesByIDs(ctx context.Context, ids []string) ([]*storage.Article, error) {
	var articles []*storage.Article

	if err := m.db.
		Where(ctx, "`id` IN (?)", ids).
		Order(ctx, "`createAt` DESC, `authorID`, `title`").
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

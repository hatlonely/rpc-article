package mysql

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-article/internal/storage"
	"gorm.io/gorm"
)

type ArticleJoinAuthor struct {
	ID       string    `gorm:"type:char(32);primary_key" json:"id"`
	AuthorID string    `gorm:"type:char(32);column:authorID;unique_index:author_title_idx;default:0" json:"authorID,omitempty"`
	Title    string    `gorm:"type:varchar(128);unique_index:author_title_idx;not null" json:"title,omitempty"`
	Tags     string    `gorm:"type:varchar(256);default:''" json:"tags,omitempty"`
	Brief    string    `gorm:"type:varchar(256);default:''" json:"brief,omitempty"`
	Content  string    `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
	CreateAt time.Time `gorm:"type:timestamp;column:createAt;not null;default:CURRENT_TIMESTAMP;index:create_at_idx" json:"createAt,omitempty"`
	UpdateAt time.Time `gorm:"type:timestamp;column:updateAt;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:update_at_idx" json:"updateAt,omitempty"`

	Key  string `gorm:"type:char(32);unique_index:key_index" json:"key,omitempty"`
	Name string `gorm:"type:char(32);not null" json:"name,omitempty"`
}

func (m *MySQL) ListArticleJoinAuthor(ctx context.Context, authorID string, offset int32, limit int32) ([]*storage.ArticleJoinAuthor, error) {
	var articleJoinAuthors []*ArticleJoinAuthor

	db := m.db.Model(ctx, &storage.Author{}).Joins(ctx, "left join authors on articles.authorID = authors.id")
	if authorID != "" {
		db = db.Where(ctx, "`authorID`=?", authorID)
	}
	if err := db.
		Order(ctx, "`createAt` DESC, `authorID`, `title`").
		Offset(ctx, int(offset)).
		Limit(ctx, int(limit)).
		Scan(ctx, &articleJoinAuthors).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var res []*storage.ArticleJoinAuthor
	for _, articleJoinAuthor := range articleJoinAuthors {
		res = append(res, &storage.ArticleJoinAuthor{
			Article: &storage.Article{
				ID:       articleJoinAuthor.ID,
				AuthorID: articleJoinAuthor.AuthorID,
				Title:    articleJoinAuthor.Title,
				Tags:     articleJoinAuthor.Tags,
				Brief:    articleJoinAuthor.Brief,
				Content:  articleJoinAuthor.Content,
				CreateAt: time.Time{},
				UpdateAt: time.Time{},
			},
			Author: &storage.Author{
				ID:   articleJoinAuthor.AuthorID,
				Key:  articleJoinAuthor.Key,
				Name: articleJoinAuthor.Name,
			},
		})
	}

	return res, nil
}

package storage

import (
	"context"
	"reflect"
	"time"

	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"
)

var ErrNotFound = errors.New("not found")

type Author struct {
	ID   string `gorm:"type:char(32);primary_key" json:"id"`
	Key  string `gorm:"type:char(32);unique_index:key_index" json:"key,omitempty"`
	Name string `gorm:"type:char(32);not null" json:"name,omitempty"`
}

type Article struct {
	ID       string    `gorm:"type:char(32);primary_key" json:"id"`
	AuthorID string    `gorm:"type:char(32);column:authorID;unique_index:author_title_idx;default:0" json:"authorID,omitempty"`
	Title    string    `gorm:"type:varchar(128);unique_index:author_title_idx;not null" json:"title,omitempty"`
	Tags     string    `gorm:"type:varchar(256);default:''" json:"tags,omitempty"`
	Brief    string    `gorm:"type:varchar(256);default:''" json:"brief,omitempty"`
	Content  string    `gorm:"type:longtext COLLATE utf8mb4_unicode_520_ci;not null" json:"content,omitempty"`
	CreateAt time.Time `gorm:"type:timestamp;column:createAt;not null;default:CURRENT_TIMESTAMP;index:create_at_idx" json:"createAt,omitempty"`
	UpdateAt time.Time `gorm:"type:timestamp;column:updateAt;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;index:update_at_idx" json:"updateAt,omitempty"`
}

type Storage interface {
	PutArticle(ctx context.Context, article *Article) (string, error)
	GetArticle(ctx context.Context, id string) (*Article, error)
	GetArticleByAuthorAndTitle(ctx context.Context, authorID string, title string) (*Article, error)
	UpdateArticle(ctx context.Context, article *Article) error
	DelArticle(ctx context.Context, id string) error
	ListArticles(ctx context.Context, offset int32, limit int32) ([]*Article, error)
	ListArticlesByAuthor(ctx context.Context, authorID string, offset int32, limit int32) ([]*Article, error)
	ListArticlesByIDs(ctx context.Context, ids []string) ([]*Article, error)

	PutAuthor(ctx context.Context, author *Author) (string, error)
	GetAuthor(ctx context.Context, id string) (*Author, error)
	UpdateAuthor(ctx context.Context, author *Author) error
	DelAuthor(ctx context.Context, id string) error
	GetAuthorByKey(ctx context.Context, key string) (*Author, error)
	UpdateAuthorByKey(ctx context.Context, author *Author) error
	DelAuthorByKey(ctx context.Context, key string) error
}

func RegisterStorage(key string, constructor interface{}) {
	refx.Register(reflect.TypeOf((*Storage)(nil)).Elem(), key, constructor)
}

func NewStorageWithOptions(options *refx.TypeOptions, opts ...refx.Option) (Storage, error) {
	v, err := refx.New(reflect.TypeOf((*Storage)(nil)).Elem(), options, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "refx.New failed")
	}
	return v.(Storage), nil
}

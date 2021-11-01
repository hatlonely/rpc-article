package article

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func NewMySQLDB(options *wrap.GORMDBWrapperOptions, opts ...refx.Option) (*MySQLDB, error) {
	db, err := wrap.NewGORMDBWrapperWithOptions(options, opts...)

	if err != nil {
		return nil, err
	}

	if !db.HasTable(&Article{}) {
		if err := db.Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").CreateTable(context.Background(), &Article{}).Unwrap().Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.AutoMigrate(context.Background(), &Article{}).Unwrap().Error; err != nil {
			return nil, err
		}
	}

	return &MySQLDB{
		db: db,
	}, nil
}

type MySQLDB struct {
	db *wrap.GORMDBWrapper
}

func (m *MySQLDB) ListArticles(ctx context.Context, offset int32, limit int32) ([]*Article, error) {
	var articles []*Article

	if err := m.db.
		Select(ctx, "id, title, tags, author_id, ctime, utime, brief").
		Order(ctx, "utime DESC").
		Offset(ctx, offset).
		Limit(ctx, limit).
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

func (m *MySQLDB) ListArticlesByAuthor(ctx context.Context, authorID int, offset, limit int) ([]*Article, error) {
	var articles []*Article

	if err := m.db.
		Select(ctx, "id, title, tags, author_id, ctime, utime, brief").
		Where(ctx, "author_id=?", authorID).
		Order(ctx, "utime DESC").
		Offset(ctx, offset).
		Limit(ctx, limit).
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

func (m *MySQLDB) ListArticlesByIDs(ctx context.Context, ids []string) ([]*Article, error) {
	var articles []*Article

	if err := m.db.
		Select(ctx, "id, title, tags, author_id, ctime, utime, brief").
		Where(ctx, "id IN (?)", ids).
		Order(ctx, "utime DESC").
		Find(ctx, &articles).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return articles, nil
}

func (m *MySQLDB) GetArticle(ctx context.Context, id string) (*Article, error) {
	var article Article
	if err := m.db.
		Where(ctx, "id=?", id).
		First(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQLDB) GetArticleByAuthorAndTitle(ctx context.Context, authorID int, title string) (*Article, error) {
	var article Article
	if err := m.db.
		Where(ctx, &Article{AuthorID: authorID, Title: title}).
		Find(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQLDB) PutArticle(ctx context.Context, article *Article) error {
	article.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return m.db.Create(ctx, article).Unwrap().Error
}

func (m *MySQLDB) UpdateArticle(ctx context.Context, article *Article) error {
	return m.db.Model(ctx, article).Where(ctx, "id=?", article.ID).Updates(ctx, article).Unwrap().Error
}

func (m *MySQLDB) DelArticle(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &Article{ID: id}).Unwrap().Error
}

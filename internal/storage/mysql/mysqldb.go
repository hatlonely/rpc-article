package article

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-article/internal/storage"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func NewMySQLDB(options *wrap.GORMDBWrapperOptions, opts ...refx.Option) (*MySQLDB, error) {
	db, err := wrap.NewGORMDBWrapperWithOptions(options, opts...)

	if err != nil {
		return nil, err
	}

	if !db.HasTable(&storage.Article{}) {
		if err := db.
			Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(context.Background(), &storage.Article{}).
			Unwrap().Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.AutoMigrate(context.Background(), &storage.Article{}).Unwrap().Error; err != nil {
			return nil, err
		}
	}

	if !db.HasTable(&storage.Author{}) {
		if err := db.
			Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(context.Background(), &storage.Author{}).
			Unwrap().Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.AutoMigrate(context.Background(), &storage.Author{}).Unwrap().Error; err != nil {
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

func (m *MySQLDB) PutAuthor(ctx context.Context, author *storage.Author) error {
	author.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return m.db.Create(ctx, author).Unwrap().Error
}

func (m *MySQLDB) GetAuthor(ctx context.Context, id string) (*storage.Author, error) {
	var author storage.Author
	if err := m.db.
		Where(ctx, "id=?", id).
		First(ctx, &author).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &author, nil
}

func (m *MySQLDB) UpdateAuthor(ctx context.Context, author *storage.Author) error {
	return m.db.Model(ctx, author).Where(ctx, "id=?", author.ID).Updates(ctx, author).Unwrap().Error
}

func (m *MySQLDB) DelAuthor(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Author{ID: id}).Unwrap().Error
}

func (m *MySQLDB) GetAuthorByKey(ctx context.Context, key string) (*storage.Author, error) {
	var author storage.Author
	if err := m.db.
		Where(ctx, "key=?", key).
		First(ctx, &author).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &author, nil
}

func (m *MySQLDB) UpdateAuthorByKey(ctx context.Context, author *storage.Author) error {
	return m.db.Model(ctx, author).Where(ctx, "key=?", author.Key).Updates(ctx, author).Unwrap().Error
}

func (m *MySQLDB) PutArticle(ctx context.Context, article *storage.Article) error {
	article.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return m.db.Create(ctx, article).Unwrap().Error
}

func (m *MySQLDB) GetArticle(ctx context.Context, id string) (*storage.Article, error) {
	var article storage.Article
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

func (m *MySQLDB) GetArticleByAuthorAndTitle(ctx context.Context, authorID int, title string) (*storage.Article, error) {
	var article storage.Article
	if err := m.db.
		Where(ctx, &storage.Article{AuthorID: authorID, Title: title}).
		Find(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQLDB) UpdateArticle(ctx context.Context, article *storage.Article) error {
	return m.db.Model(ctx, article).Where(ctx, "id=?", article.ID).Updates(ctx, article).Unwrap().Error
}

func (m *MySQLDB) DelArticle(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Article{ID: id}).Unwrap().Error
}

func (m *MySQLDB) ListArticles(ctx context.Context, offset int32, limit int32) ([]*storage.Article, error) {
	var articles []*storage.Article

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

func (m *MySQLDB) ListArticlesByAuthor(ctx context.Context, authorID int, offset, limit int) ([]*storage.Article, error) {
	var articles []*storage.Article

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

func (m *MySQLDB) ListArticlesByIDs(ctx context.Context, ids []string) ([]*storage.Article, error) {
	var articles []*storage.Article

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

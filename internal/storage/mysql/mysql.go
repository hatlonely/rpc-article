package mysql

import (
	"context"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-article/internal/storage"
)

func NewMySQL(options *wrap.GORMDBWrapperOptions, opts ...refx.Option) (*MySQL, error) {
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

	return &MySQL{
		db: db,
	}, nil
}

type MySQL struct {
	db *wrap.GORMDBWrapper
}

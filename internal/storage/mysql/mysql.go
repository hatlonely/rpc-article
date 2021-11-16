package mysql

import (
	"context"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-article/internal/storage"
)

func init() {
	storage.RegisterStorage("MySQL", NewMySQLWithOptions)
}

func NewMySQLWithOptions(options *wrap.GORMDBWrapperOptions, opts ...refx.Option) (*MySQL, error) {
	db, err := wrap.NewGORMDBWrapperWithOptions(options, opts...)

	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(context.Background(), &storage.Article{}); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(context.Background(), &storage.Author{}); err != nil {
		return nil, err
	}

	return &MySQL{
		db: db,
	}, nil
}

type MySQL struct {
	db *wrap.GORMDBWrapper
}

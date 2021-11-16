package mysql

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/rpc-article/internal/storage"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (m *MySQL) PutAuthor(ctx context.Context, author *storage.Author) (string, error) {
	author.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return author.ID, m.db.Create(ctx, author).Unwrap().Error
}

func (m *MySQL) GetAuthor(ctx context.Context, id string) (*storage.Author, error) {
	var author storage.Author
	if err := m.db.
		Where(ctx, "`id`=?", id).
		First(ctx, &author).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, storage.ErrNotFound
		}

		return nil, err
	}

	return &author, nil
}

func (m *MySQL) UpdateAuthor(ctx context.Context, author *storage.Author) error {
	return m.db.Model(ctx, author).Where(ctx, "`id`=?", author.ID).Updates(ctx, author).Unwrap().Error
}

func (m *MySQL) DelAuthor(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Author{ID: id}).Unwrap().Error
}

func (m *MySQL) GetAuthorByKey(ctx context.Context, key string) (*storage.Author, error) {
	var author storage.Author
	if err := m.db.
		Where(ctx, "`key`=?", key).
		First(ctx, &author).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, storage.ErrNotFound
		}

		return nil, err
	}

	return &author, nil
}

func (m *MySQL) UpdateAuthorByKey(ctx context.Context, author *storage.Author) error {
	return m.db.Model(ctx, author).Where(ctx, "`key`=?", author.Key).Updates(ctx, author).Unwrap().Error
}

func (m *MySQL) DelAuthorByKey(ctx context.Context, key string) error {
	return m.db.Delete(ctx, &storage.Author{Key: key}).Unwrap().Error
}

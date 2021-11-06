package mysql

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hatlonely/rpc-article/internal/storage"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
)

var username = "hatlonely"
var password = "keaiduo1"
var database = "article"
var host = "127.0.0.1"
var port = 3306

func TestNewMySQLWithOptions(t *testing.T) {
	Convey("TestNewMySQLWithOptions", t, func() {
		db, err := NewMySQLWithOptions(&wrap.GORMDBWrapperOptions{
			Retry: micro.RetryOptions{
				Attempts:      1,
				LastErrorOnly: true,
			},
			Wrapper: wrap.WrapperOptions{
				Name: "articleDB",
			},
			Gorm: wrap.GormOptions{
				Username:        username,
				Password:        password,
				Database:        database,
				Host:            host,
				Port:            port,
				ConnMaxLifeTime: 60 * time.Second,
				MaxIdleConns:    10,
				MaxOpenConns:    20,
				LogMode:         false,
			},
		})
		So(err, ShouldBeNil)
		So(db, ShouldNotBeNil)
	})
}

func NewTestMysql() (*MySQL, error) {
	return NewMySQLWithOptions(&wrap.GORMDBWrapperOptions{
		Retry: micro.RetryOptions{
			Attempts:      1,
			LastErrorOnly: true,
		},
		Wrapper: wrap.WrapperOptions{
			Name: "articleDB",
		},
		Gorm: wrap.GormOptions{
			Username:        username,
			Password:        password,
			Database:        database,
			Host:            host,
			Port:            port,
			ConnMaxLifeTime: 60 * time.Second,
			MaxIdleConns:    10,
			MaxOpenConns:    20,
			LogMode:         false,
		},
	})
}

func CleanTestMysql(db *MySQL) {
	db.db.Delete(context.Background(), &storage.Author{Key: "testKey1"})
	db.db.Delete(context.Background(), &storage.Article{AuthorID: "testAuthorID1", Title: "testTitle1"})
	for i := 0; i < 10; i++ {
		db.db.Delete(context.Background(), &storage.Article{
			AuthorID: fmt.Sprintf("testAuthorID%v", i),
			Title:    fmt.Sprintf("testTitle%v", i),
		})

		db.db.Delete(context.Background(), &storage.Article{
			AuthorID: "testAuthorID0",
			Title:    fmt.Sprintf("testTitle%v", i),
		})
	}
}

package mysql

import (
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"
	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
)

var username = "hatlonely"
var password = "keaiduo1"
var database = "article"
var host = "127.0.0.1"
var port = 3306

func TestNewMySQLDB(t *testing.T) {
	Convey("TestNewMySQLDB", t, func() {
		db, err := NewMySQL(&wrap.GORMDBWrapperOptions{
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

func TestMySQLDB_PutArticle(t *testing.T) {
	Convey("TestMySQLDB_PutArticle", t, func() {

	})
}

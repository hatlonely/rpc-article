package article

import (
	"testing"
	"time"

	"github.com/hatlonely/go-kit/micro"

	"github.com/hatlonely/go-kit/wrap"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMySQLDB(t *testing.T) {
	Convey("TestNewMySQLDB", t, func() {
		db, err := NewMySQLDB(&wrap.GORMDBWrapperOptions{
			Retry: micro.RetryOptions{
				Attempts:      1,
				LastErrorOnly: true,
			},
			Wrapper: wrap.WrapperOptions{
				Name: "articleDB",
			},
			Gorm: wrap.GormOptions{
				Username:        "hatlonely",
				Password:        "keaiduo1",
				Database:        "article",
				Host:            "127.0.0.1",
				Port:            3306,
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

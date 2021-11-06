package mysql

import (
	"context"
	"testing"

	"github.com/hatlonely/rpc-article/internal/storage"

	"github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMySQLDB_PutAuthor(t *testing.T) {
	Convey("TestMySQLDB_PutAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)
		})

		Convey("duplicate entry", func() {
			err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)

			err = db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldNotBeNil)
			So(err.(*mysql.MySQLError).Number, ShouldEqual, 1062)
			So(err.(*mysql.MySQLError).Message, ShouldContainSubstring, "Duplicate entry")
		})
	})
}

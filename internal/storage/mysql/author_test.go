package mysql

import (
	"context"
	"testing"

	"github.com/hatlonely/rpc-article/internal/storage"

	"github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMySQL_PutAuthor(t *testing.T) {
	Convey("TestMySQL_PutAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			_, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)
		})

		Convey("duplicate entry", func() {
			_, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)

			_, err = db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldNotBeNil)
			So(err.(*mysql.MySQLError).Number, ShouldEqual, 1062)
			So(err.(*mysql.MySQLError).Message, ShouldContainSubstring, "Duplicate entry")
		})
	})
}

func TestMySQL_GetAuthor(t *testing.T) {
	Convey("TestMySQL_GetAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:   id,
				Key:  "testKey1",
				Name: "testName1",
			})
		})

		Convey("not exist author", func() {
			author, err := db.GetAuthor(context.Background(), "NotExistID")
			So(err, ShouldBeNil)
			So(author, ShouldBeNil)
		})
	})
}

func TestMySQL_UpdateAuthor(t *testing.T) {
	Convey("TestMySQL_UpdateAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)

			err = db.UpdateAuthor(context.Background(), &storage.Author{
				ID:   id,
				Name: "testName2",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:   id,
				Key:  "testKey1",
				Name: "testName2",
			})
		})
	})
}

func TestMySQL_DelAuthor(t *testing.T) {
	Convey("TestMySQL_DelAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:  "testKey1",
				Name: "testName1",
			})
			So(err, ShouldBeNil)

			err = db.DelAuthor(context.Background(), id)
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldBeNil)
		})
	})
}

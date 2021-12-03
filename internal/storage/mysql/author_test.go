package mysql

import (
	"context"
	"testing"

	"github.com/hatlonely/rpc-article/internal/storage"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMySQL_PutAuthor(t *testing.T) {
	Convey("TestMySQL_PutAuthor", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			_, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)
		})

		Convey("duplicate entry", func() {
			id1, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			id2, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName2",
				Avatar: "testAvatar2",
			})
			So(err, ShouldBeNil)
			So(id1, ShouldEqual, id2)

			author, err := db.GetAuthor(context.Background(), id1)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:     id1,
				Key:    "testKey1",
				Name:   "testName2",
				Avatar: "testAvatar2",
			})
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
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:     id,
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
		})

		Convey("not exist author", func() {
			_, err := db.GetAuthor(context.Background(), "NotExistID")
			So(err, ShouldEqual, storage.ErrNotFound)
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
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			err = db.UpdateAuthor(context.Background(), &storage.Author{
				ID:     id,
				Name:   "testName2",
				Avatar: "testAvatar2",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:     id,
				Key:    "testKey1",
				Name:   "testName2",
				Avatar: "testAvatar2",
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
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			err = db.DelAuthor(context.Background(), id)
			So(err, ShouldBeNil)

			_, err = db.GetAuthor(context.Background(), id)
			So(err, ShouldEqual, storage.ErrNotFound)
		})
	})
}

func TestMySQL_GetAuthorByKey(t *testing.T) {
	Convey("TestMySQL_GetAuthorByKey", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthorByKey(context.Background(), "testKey1")
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:     id,
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
		})

		Convey("not exist author", func() {
			_, err := db.GetAuthorByKey(context.Background(), "NotExistKey")
			So(err, ShouldEqual, storage.ErrNotFound)
		})
	})
}

func TestMySQL_UpdateAuthorByKey(t *testing.T) {
	Convey("TestMySQL_UpdateAuthorByKey", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			err = db.UpdateAuthorByKey(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName2",
				Avatar: "testAvatar2",
			})
			So(err, ShouldBeNil)

			author, err := db.GetAuthor(context.Background(), id)
			So(err, ShouldBeNil)
			So(author, ShouldResemble, &storage.Author{
				ID:     id,
				Key:    "testKey1",
				Name:   "testName2",
				Avatar: "testAvatar2",
			})
		})
	})
}

func TestMySQL_DelAuthorByKey(t *testing.T) {
	Convey("TestMySQL_DelAuthorByKey", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutAuthor(context.Background(), &storage.Author{
				Key:    "testKey1",
				Name:   "testName1",
				Avatar: "testAvatar1",
			})
			So(err, ShouldBeNil)

			err = db.DelAuthorByKey(context.Background(), "testKey1")
			So(err, ShouldBeNil)

			_, err = db.GetAuthor(context.Background(), id)
			So(err, ShouldEqual, storage.ErrNotFound)
		})
	})
}

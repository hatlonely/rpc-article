package mysql

import (
	"context"
	"testing"

	"github.com/hatlonely/rpc-article/internal/storage"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMySQL_PutArticle(t *testing.T) {
	Convey("TestMySQL_PutArticle", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			_, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID",
				Title:    "testTitle",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief",
				Content:  "testContent",
			})
			So(err, ShouldBeNil)
		})

		//Convey("duplicate entry", func() {
		//	_, err := db.PutAuthor(context.Background(), &storage.Author{
		//		Key:  "testKey1",
		//		Name: "testName1",
		//	})
		//	So(err, ShouldBeNil)
		//
		//	_, err = db.PutAuthor(context.Background(), &storage.Author{
		//		Key:  "testKey1",
		//		Name: "testName1",
		//	})
		//	So(err, ShouldNotBeNil)
		//	So(err.(*mysql.MySQLError).Number, ShouldEqual, 1062)
		//	So(err.(*mysql.MySQLError).Message, ShouldContainSubstring, "Duplicate entry")
		//})
	})
}

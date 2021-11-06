package mysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-sql-driver/mysql"
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
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)
		})

		Convey("duplicate entry", func() {
			_, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			_, err = db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldNotBeNil)
			So(err.(*mysql.MySQLError).Number, ShouldEqual, 1062)
			So(err.(*mysql.MySQLError).Message, ShouldContainSubstring, "Duplicate entry")
		})
	})
}

func TestMySQL_GetArticle(t *testing.T) {
	Convey("TestMySQL_GetArticle", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			article, err := db.GetArticle(context.Background(), id)
			So(err, ShouldBeNil)
			So(article, ShouldResemble, &storage.Article{
				ID:       id,
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
			})
		})

		Convey("not exist article", func() {
			article, err := db.GetArticle(context.Background(), "NotExistID")
			So(err, ShouldBeNil)
			So(article, ShouldBeNil)
		})
	})
}

func TestMySQL_UpdateArticle(t *testing.T) {
	Convey("TestMySQL_UpdateArticle", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			err = db.UpdateArticle(context.Background(), &storage.Article{
				ID:       id,
				AuthorID: "testAuthorID2",
				Title:    "testTitle2",
				Tags:     "testTag3,testTag4",
				Brief:    "testBrief2",
				Content:  "testContent2",
			})

			article, err := db.GetArticle(context.Background(), id)
			So(err, ShouldBeNil)
			So(article, ShouldResemble, &storage.Article{
				ID:       id,
				AuthorID: "testAuthorID2",
				Title:    "testTitle2",
				Tags:     "testTag3,testTag4",
				Brief:    "testBrief2",
				Content:  "testContent2",
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
			})
		})
	})
}

func TestMySQL_DelArticle(t *testing.T) {
	Convey("TestMySQL_DelArticle", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			err = db.DelArticle(context.Background(), id)
			So(err, ShouldBeNil)

			article, err := db.GetArticle(context.Background(), id)
			So(err, ShouldBeNil)
			So(article, ShouldBeNil)
		})
	})
}

func TestMySQL_GetArticleByAuthorAndTitle(t *testing.T) {
	Convey("TestMySQL_GetArticleByAuthorAndTitle", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			id, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			article, err := db.GetArticleByAuthorAndTitle(context.Background(), "testAuthorID1", "testTitle1")
			So(err, ShouldBeNil)
			So(article, ShouldResemble, &storage.Article{
				ID:       id,
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
			})
		})
	})
}

func TestMySQL_ListArticles(t *testing.T) {
	Convey("TestMySQL_ListArticles", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		for i := 0; i < 10; i++ {
			_, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: fmt.Sprintf("testAuthorID%v", i),
				Title:    fmt.Sprintf("testTitle%v", i),
				Tags:     fmt.Sprintf("testTag%v", i),
				Brief:    fmt.Sprintf("testBrief%v", i),
				Content:  fmt.Sprintf("testContent%v", i),
			})
			So(err, ShouldBeNil)
		}

		Convey("normal", func() {
			articles, err := db.ListArticles(context.Background(), 0, 10)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 10)

			for i := 0; i < 10; i++ {
				So(articles[i], ShouldResemble, &storage.Article{
					ID:       articles[i].ID,
					AuthorID: fmt.Sprintf("testAuthorID%v", i),
					Title:    fmt.Sprintf("testTitle%v", i),
					Tags:     fmt.Sprintf("testTag%v", i),
					Brief:    fmt.Sprintf("testBrief%v", i),
					Content:  fmt.Sprintf("testContent%v", i),
					CreateAt: articles[i].CreateAt,
					UpdateAt: articles[i].UpdateAt,
				})
			}
		})
	})
}

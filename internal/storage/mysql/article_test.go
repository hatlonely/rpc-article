package mysql

import (
	"context"
	"fmt"
	"sort"
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
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)
		})

		Convey("duplicate entry", func() {
			id1, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag1,testTag2",
				Brief:    "testBrief1",
				Content:  "testContent1",
			})
			So(err, ShouldBeNil)

			id2, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag3,testTag4",
				Brief:    "testBrief2",
				Content:  "testContent2",
			})
			So(err, ShouldBeNil)
			So(id1, ShouldEqual, id2)

			article, err := db.GetArticle(context.Background(), id1)
			So(err, ShouldBeNil)
			So(article, ShouldResemble, &storage.Article{
				ID:       id1,
				AuthorID: "testAuthorID1",
				Title:    "testTitle1",
				Tags:     "testTag3,testTag4",
				Brief:    "testBrief2",
				Content:  "testContent2",
				CreateAt: article.CreateAt,
				UpdateAt: article.UpdateAt,
			})
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
			_, err := db.GetArticle(context.Background(), "NotExistID")
			So(err, ShouldEqual, storage.ErrNotFound)
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

			_, err = db.GetArticle(context.Background(), id)
			So(err, ShouldEqual, storage.ErrNotFound)
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
			articles, err := db.ListArticles(context.Background(), 0, 100)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 10)

			sort.Slice(articles, func(i, j int) bool {
				return articles[i].AuthorID < articles[j].AuthorID
			})
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

		Convey("normal offset limit", func() {
			articles, err := db.ListArticles(context.Background(), 3, 4)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 4)
		})
	})
}

func TestMySQL_ListArticlesByAuthor(t *testing.T) {
	Convey("TestMySQL_ListArticles", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		for i := 0; i < 10; i++ {
			_, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: "testAuthorID0",
				Title:    fmt.Sprintf("testTitle%v", i),
				Tags:     fmt.Sprintf("testTag%v", i),
				Brief:    fmt.Sprintf("testBrief%v", i),
				Content:  fmt.Sprintf("testContent%v", i),
			})
			So(err, ShouldBeNil)
		}

		Convey("normal", func() {
			articles, err := db.ListArticlesByAuthor(context.Background(), "testAuthorID0", 0, 100)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 10)

			sort.Slice(articles, func(i, j int) bool {
				return articles[i].Title < articles[j].Title
			})
			for i := 0; i < 10; i++ {
				So(articles[i], ShouldResemble, &storage.Article{
					ID:       articles[i].ID,
					AuthorID: "testAuthorID0",
					Title:    fmt.Sprintf("testTitle%v", i),
					Tags:     fmt.Sprintf("testTag%v", i),
					Brief:    fmt.Sprintf("testBrief%v", i),
					Content:  fmt.Sprintf("testContent%v", i),
					CreateAt: articles[i].CreateAt,
					UpdateAt: articles[i].UpdateAt,
				})
			}
		})

		Convey("normal offset limit", func() {
			articles, err := db.ListArticlesByAuthor(context.Background(), "testAuthorID0", 3, 4)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 4)
		})
	})
}

func TestMySQL_ListArticlesByIDs(t *testing.T) {
	Convey("TestMySQL_ListArticlesByIDs", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		Convey("normal", func() {
			var ids []string
			for i := 0; i < 10; i++ {
				id, err := db.PutArticle(context.Background(), &storage.Article{
					AuthorID: fmt.Sprintf("testAuthorID%v", i),
					Title:    fmt.Sprintf("testTitle%v", i),
					Tags:     fmt.Sprintf("testTag%v", i),
					Brief:    fmt.Sprintf("testBrief%v", i),
					Content:  fmt.Sprintf("testContent%v", i),
				})
				So(err, ShouldBeNil)
				ids = append(ids, id)
			}

			articles, err := db.ListArticlesByIDs(context.Background(), ids)
			So(err, ShouldBeNil)
			So(len(articles), ShouldEqual, 10)

			sort.Slice(articles, func(i, j int) bool {
				return articles[i].AuthorID < articles[j].AuthorID
			})
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

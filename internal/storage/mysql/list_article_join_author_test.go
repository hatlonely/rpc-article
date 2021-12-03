package mysql

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/hatlonely/rpc-article/internal/storage"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMySQL_ListArticleJoinAuthor(t *testing.T) {
	Convey("TestMySQL_ListArticles", t, func() {
		db, err := NewTestMysql()
		So(err, ShouldBeNil)

		CleanTestMysql(db)

		authorID, err := db.PutAuthor(context.Background(), &storage.Author{
			Key:    "testKey1",
			Name:   "testName1",
			Avatar: "testAvatar1",
		})
		So(err, ShouldBeNil)
		for i := 0; i < 10; i++ {
			_, err := db.PutArticle(context.Background(), &storage.Article{
				AuthorID: authorID,
				Title:    fmt.Sprintf("testTitle%v", i),
				Tags:     fmt.Sprintf("testTag%v", i),
				Brief:    fmt.Sprintf("testBrief%v", i),
				Content:  fmt.Sprintf("testContent%v", i),
			})
			So(err, ShouldBeNil)
		}

		Convey("normal", func() {
			articleJoinAuthors, err := db.ListArticleJoinAuthor(context.Background(), authorID, 0, 100)
			So(err, ShouldBeNil)
			So(len(articleJoinAuthors), ShouldEqual, 10)

			sort.Slice(articleJoinAuthors, func(i, j int) bool {
				return articleJoinAuthors[i].Article.Title < articleJoinAuthors[j].Article.Title
			})
			for i := 0; i < 10; i++ {
				So(articleJoinAuthors[i], ShouldResemble, &storage.ArticleJoinAuthor{
					Article: &storage.Article{
						ID:       articleJoinAuthors[i].Article.ID,
						AuthorID: authorID,
						Title:    fmt.Sprintf("testTitle%v", i),
						Tags:     fmt.Sprintf("testTag%v", i),
						Brief:    fmt.Sprintf("testBrief%v", i),
						Content:  fmt.Sprintf("testContent%v", i),
						CreateAt: articleJoinAuthors[i].Article.CreateAt,
						UpdateAt: articleJoinAuthors[i].Article.UpdateAt,
					},
					Author: &storage.Author{
						ID:     authorID,
						Key:    "testKey1",
						Name:   "testName1",
						Avatar: "testAvatar1",
					},
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

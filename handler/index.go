/*
前端页面
handler
*/

package handler

import (
	"github.com/zituocn/gow"
)

// Index 首页
// route /
func Index(c *gow.Context) {
	c.HTML("index.html")
}

// Category 分类目录
//	route /category
func Category(c *gow.Context) {

}

// ArticleList 分类文章列表
//	route /category/golang
func ArticleList(c *gow.Context) {
	c.HTML("article_list.html")
}

// ArticleDetail  文章详情
//	route /article-{id}.html
func ArticleDetail(c *gow.Context) {
	c.HTML("article_detail.html")
}

// Archive
//	route /archive
func Archive(c *gow.Context) {
	c.HTML("archive.html")
}

// ArchiveMonth
//	route /archive/202011
func ArchiveMonth(c *gow.Context) {
	c.HTML("archive_month.html")
}

// User
//	route /user
func User(c *gow.Context) {
	c.HTML("user.html")
}

// ArchiveUser
//	router /user/1
func ArchiveUser(c *gow.Context) {
	c.HTML("archive_user.html")
}

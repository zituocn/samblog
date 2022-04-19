/*
admin page
*/

package admin

import "github.com/zituocn/gow"

// Login 登录
func Login(c *gow.Context) {
	c.HTML("admin/login.html")
}

// Logout 退出
func Logout(c *gow.Context) {

}

// Index 管理后台首页
func Index(c *gow.Context) {
	c.HTML("admin/index.html")
}

// Article 文章列表
func Article(c *gow.Context) {
	c.HTML("admin/article.html")
}

// ArticleAdd 添加文章
func ArticleAdd(c *gow.Context) {
	c.HTML("admin/article_add.html")
}

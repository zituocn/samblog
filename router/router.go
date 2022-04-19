/*



/article-{id}.html  # 文章详情

/category  # 所有文章列表
/category/golang  # 分类文章列表

/archive   # 所有月份列表
/archive/202012  # 某个月份的文章列表

/user   # 所有用户列表
/user/1  # 某个用户的文章列表



*/

package router

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/samblog/handler"
	"github.com/zituocn/samblog/handler/admin"
	"github.com/zituocn/samblog/middleware"
)

// PageRouter page router
func PageRouter(r *gow.Engine) {

	r.Use(middleware.Version(), middleware.UserAuth())

	r.Any("/", handler.Index)
	r.Any("/category", handler.Category)
	r.Any("/category/{name}", handler.ArticleList)
	r.Any("/article-{id}.html", handler.ArticleDetail)
	r.Any("/archive", handler.Archive)
	r.Any("/archive/{month}", handler.ArchiveMonth)
	r.Any("/user", handler.User)
	r.Any("/user/{uid}", handler.ArchiveUser)

	r.GET("/login", admin.Login)
	r.GET("/logout", admin.Logout)

	//admin

	dash := r.Group("/dashboard")
	{
		dash.GET("/", admin.Index)
	}
}

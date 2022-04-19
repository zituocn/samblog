package middleware

import (
	"github.com/zituocn/gow/lib/logy"
	"github.com/zituocn/samblog/cache"
	"time"

	"github.com/zituocn/gow"
)

var (
	ver int
)

func init() {
	ver = int(time.Now().Unix())
}

// Version 版本号相关的middleware
//	输出编号等静态信息到html页面
func Version() gow.HandlerFunc {
	return func(c *gow.Context) {

		blog,err:= new(cache.BlogCache).GetBlog()
		if err!=nil{
			logy.Errorf("从缓存中获取 blog 信息失败 : %v",err)
		}
		c.Data["blog"] =blog
		c.Data["ver"] = ver
		c.Data["year"] = time.Now().Year()
		c.Next()
	}
}

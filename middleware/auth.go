package middleware

import "github.com/zituocn/gow"

// UserAuth 用户权限
func UserAuth() gow.HandlerFunc {
	return func(c *gow.Context) {
		user := GetContextUser(c)
		c.Data["auth"] = user
		c.Next()
	}
}

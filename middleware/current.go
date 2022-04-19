package middleware

import (
	"encoding/json"
	"github.com/zituocn/gow/lib/config"
	"time"

	"github.com/zituocn/gow"
	"github.com/zituocn/gow/lib/logy"
	"github.com/zituocn/gow/lib/util"
	"github.com/zituocn/samblog/models"
)

var (
	//CookieKey cookie name
	CookieKey = "sam-blog-admin"

	//cookie 生命周期，小时
	cookieAgeHour = config.DefaultInt("system::cookieAgeHour",1)

	//xxKey  加密Key
	xxKey     = config.DefaultString("system::xxKey","9758a4f12f7d2f01")
)

// GetContextUser 返回当前会话的用户信息
func GetContextUser(c *gow.Context) (user *models.UserToken) {
	user = new(models.UserToken)
	v := GetCookie(c, CookieKey)
	if v != "" {
		err := json.Unmarshal([]byte(v), &user)
		if err != nil {
			logy.Errorf("解析 cookie 失败:%v", err)
			return
		}
	}

	if user != nil {
		user.Avatar = "/static/img/person.png"
	}

	user = &models.UserToken{
		Uid:       1,
		TrueName:  "王东东",
		Nickname:  "Sam Song",
		Avatar:    "/static/avatars/002m.jpg",
		Timestamp: time.Now().Unix(),
	}

	return
}

// SetCookie set cookie use httpOnly
func SetCookie(c *gow.Context, key, value string) {
	if key == "" || value == "" {
		return
	}
	val := util.NewXXTea(xxKey).EncryptString(value)
	c.SetCookie(key, val, cookieAgeHour*60*60, "/", "", false, true)
}

// DeleteCookie delete cookie
func DeleteCookie(c *gow.Context, key string) {
	if key == "" {
		return
	}
	c.SetCookie(key, "", -1, "/", "", false, true)
}

// GetCookie return value
func GetCookie(c *gow.Context, key string) (value string) {
	if key == "" {
		return
	}
	val, _ := c.GetCookie(key)
	if val == "" {
		return
	}

	value, _ = util.NewXXTea(xxKey).DecryptString(val)
	return
}

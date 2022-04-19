package service

import (
	"errors"

	"github.com/zituocn/gow/lib/util"
	"github.com/zituocn/samblog/models"
)

type UserLoginService struct{}

// UserLogin 登录
func (m *UserLoginService) UserLogin(username, password string) (user *models.User, err error) {
	user = new(models.User)

	err = user.GetUserByUsername(username)
	if err != nil || user.Uid == 0 {
		err = errors.New("用户名不存在")
		return
	}
	if user.Password != util.MD5(password) {
		err = errors.New("密码错误")
		return
	}

	return
}

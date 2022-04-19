package models

import (
	"github.com/zituocn/gow/lib/mysql"
	"github.com/zituocn/samblog/enum"
)

// User 用户
type User struct {
	Uid       int64            `json:"uid" gorm:"primary_key;column:uid"`
	Username  string           `json:"username" gorm:"column:username"`
	Mobile    string           `json:"mobile" gorm:"column:mobile"`
	TrueName  string           `json:"true_name" gorm:"column:true_name"`
	Avatar    string           `json:"avatar" gorm:"column:avatar"`
	Password  string           `json:"password" gorm:"column:password"`
	Sex       enum.SexType     `json:"sex" gorm:"column:sex"`
	Count     int              `json:"count" gorm:"column:count"`
	UserState enum.CommonState `json:"user_state" gorm:"column:user_state"`
	DateTime
}

// TableName 表名
func (*User) TableName() string {
	return "tbl_user"
}

// GetUserByUsername 取用户
func (m *User) GetUserByUsername(username string) error {
	db := mysql.GetORM()
	return db.Model(m).Where("username=?", username).First(&m).Error
}

// GetUserList 所有用户的列表
func (m *User) GetUserList() (data []*User, err error) {
	db := mysql.GetORM()
	err = db.Model(m).Find(&data).Error
	return
}

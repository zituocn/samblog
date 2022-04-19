package models

import "github.com/zituocn/gow/lib/mysql"

// Blog 信息
type Blog struct {
	Id          int    `gorm:"primary_key;column:id" json:"id"`       //id
	Name        string `gorm:"column:name" json:"name"`               //名称
	Keyword     string `gorm:"column:keyword" json:"keyword"`         //关键字
	Description string `gorm:"column:description" json:"description"` //介绍
	RecordNum   string `gorm:"column:record_num" json:"record_num"`   //备案号
	DateTime
}

func (*Blog) TableName() string {
	return "tbl_blog"
}

// GetBlog 取配置信息
func (m *Blog) GetBlog() (blog *Blog ,err error) {
	db := mysql.GetORM()
	err=  db.Model(m).First(&blog).Error
	return
}

// Update Update
func (m *Blog) Update(mp map[string]interface{}) error {
	db := mysql.GetORM()
	return db.Model(m).Updates(mp).Error
}

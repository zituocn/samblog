package models

import "github.com/zituocn/gow/lib/mysql"

// Month 月份
type Month struct {
	Id    int64 `gorm:"primary_key;column:id" json:"id"`
	Month int   `gorm:"column:month" json:"month"`
	Count int   `gorm:"column:count" json:"count"`
	DateTime
}

func (*Month) TableName() string {
	return "tbl_month"
}

// GetMonthList 获取所有的月份信息
func (m *Month) GetMonthList() (data []*Month, err error) {
	db := mysql.GetORM()
	err = db.Model(m).Find(&data).Error
	return
}

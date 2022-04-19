package cache

import (
	"strings"

	"github.com/zituocn/gow/lib/util"
	"github.com/zituocn/samblog/models"
)

type CategoryCache struct{}

var (
	categoryKey = "category"
	mc          *util.MemoryCache
)

func init() {
	mc = util.NewMemoryCache(24, 1)
}

func (m *CategoryCache) Remove()  error {
	key := categoryKey
	return  mc.RemoveCache(key)
}

// GetAll 缓存中的所有数据
func (m *CategoryCache) GetAll() (data []*models.Category, err error) {
	key := categoryKey
	ok, err := mc.GetCache(key, &data)
	if !ok {
		data, err := new(models.Category).GetCategoryList()
		if err != nil {
			return
		}
		mc.SetCache(key, data)
	}
	return
}

// GetCategoryByCid 根据cid取category信息
func (m *CategoryCache) GetCategoryByCid(cid int64) (category *models.Category) {
	data, err := m.GetAll()
	if err != nil {
		return
	}
	for _, item := range data {
		if item.Cid == cid {
			category = item
			return
		}
	}
	return
}

// GetCategoryByEName 根据名称取category信息
func (m *CategoryCache) GetCategoryByEName(eName string) (category *models.Category) {
	data, err := m.GetAll()
	if err != nil {
		return
	}
	for _, item := range data {
		if strings.ToLower(item.EName) == strings.ToLower(eName) {
			category = item
			return
		}
	}
	return
}

package cache

import "github.com/zituocn/samblog/models"

type BlogCache struct{}

var(

	// blogKey 缓存key
	blogKey = "blog"
)

// GetBlog 缓存中的所有数据
func (m *BlogCache) GetBlog() (data *models.Blog, err error) {
	key := blogKey
	ok, err := mc.GetCache(key, &data)
	if !ok {
		data, err := new(models.Blog).GetBlog()
		if err != nil {
			return
		}
		mc.SetCache(key, data)
	}
	return
}

// Remove 清理缓存
func (m *BlogCache) Remove()  error {
	key := blogKey
	return  mc.RemoveCache(key)
}
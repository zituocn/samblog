package models

// UserToken 用户后台auth
type UserToken struct {
	Uid       int64  `json:"uid"`
	TrueName  string `json:"true_name"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Timestamp int64  `json:"timestamp"`
	IsLock    bool   `json:"is_lock"`
}

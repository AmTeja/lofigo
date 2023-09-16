package models

type Session struct {
	Id      uint  `json:"id"`
	UserID  uint  `json:"user_id"`
	Expires int64 `json:"expires"`
}

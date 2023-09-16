package models

type Post struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category_id"`
	SubCategory string `json:"sub_category_id"`
	Attachments string `json:"attachments"`
	UserId      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserId" json:"user"`
	CreatedAt   uint64 `json:"created_at"`
	UpdatedAt   uint64 `json:"updated_at"`
}

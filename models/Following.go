package models

type Following struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	UserId      uint `json:"user_id" gorm:"not null,index,uniqueIndex:idx_user_id_following_id,foreignKey:UserId,references:users,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FollowingId uint `json:"following_id" gorm:"not null,index,uniqueIndex:idx_user_id_following_id,foreignKey:FollowingId,references:users,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

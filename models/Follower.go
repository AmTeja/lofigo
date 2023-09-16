package models

type Follower struct {
	Id         uint `json:"id" gorm:"primaryKey"`
	UserId     uint `json:"user_id" gorm:"not null,index,uniqueIndex:idx_user_id_follower_id,foreignKey:UserId,references:users,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FollowerId uint `json:"follower_id" gorm:"not null,index,uniqueIndex:idx_user_id_follower_id,foreignKey:FollowerId,references:users,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

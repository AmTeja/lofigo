package models

type Key struct { // add  foreign key constraint on UserID
	Id             string `json:"id" gorm:"primaryKey"`
	HashedPassword string `json:"hashed_password" gorm:"not null"`
	UserID         uint   `json:"user_id" gorm:"not null,index,uniqueIndex:idx_user_id_key_id,foreignKey:UserID,references:users,constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	IsPrimary      bool   `json:"primary" gorm:"not null"`
}

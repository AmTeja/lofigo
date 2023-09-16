package models

//Uses mysql

type User struct {
	Id                uint        `json:"id"`
	Name              string      `json:"name"`
	Email             string      `json:"email" gorm:"unique"`
	FirstName         string      `json:"first_name"`
	LastName          string      `json:"last_name"`
	Followers         []Follower  `json:"followers"`
	Following         []Following `json:"following"`
	FollowersCount    int         `json:"followers_count"`
	FollowingCount    int         `json:"following_count"`
	CoverPictureUrl   string      `json:"cover_picture_url"`
	ProfilePictureUrl string      `json:"profile_picture_url"`
	Country           string      `json:"country"`
	MainCategory      string      `json:"main_category"`
	Keys              []Key       `gorm:"foreignKey:UserID" json:"-"`
	Sessions          []Session   `gorm:"foreignKey:UserID" json:"-"`
}

// Copy With Methods
func (u User) Copy() User {
	return User{
		Id:                u.Id,
		Name:              u.Name,
		Email:             u.Email,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Followers:         u.Followers,
		Following:         u.Following,
		FollowersCount:    u.FollowersCount,
		FollowingCount:    u.FollowingCount,
		CoverPictureUrl:   u.CoverPictureUrl,
		ProfilePictureUrl: u.ProfilePictureUrl,
		Country:           u.Country,
		MainCategory:      u.MainCategory,
		Keys:              u.Keys,
		Sessions:          u.Sessions,
	}
}

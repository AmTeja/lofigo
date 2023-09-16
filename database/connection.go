package database

import (
	"github.com/amteja/lofig/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the database connection
var DB *gorm.DB

// Connect opens a database connection
func Connect() {
	database, err := gorm.Open(mysql.Open("root:mysql@/lofig"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = database

	migrate()
}

func migrate() {
	DB.AutoMigrate(
		&models.User{}, &models.Key{},
		&models.Session{}, &models.Post{},
		&models.Follower{}, &models.Following{})
}

package util

import (
	"github.com/amteja/lofig/database"
	"github.com/amteja/lofig/models"
	"github.com/golang-jwt/jwt/v5"
)

func UserFromJWT(jwToken string) (models.User, error) {
	var token, _ = jwt.Parse(jwToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	var user models.User

	var id, err = token.Claims.GetIssuer()

	if err != nil {
		return models.User{}, err
	}

	database.DB.Find(&user, id)

	if user.Id == 0 {
		return models.User{}, err
	}

	return user, nil
}

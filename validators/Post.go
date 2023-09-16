package validators

import (
	"errors"

	"github.com/amteja/lofig/models"
)

func ValidatePost(post models.Post) error {
	if post.Title == "" {
		return errors.New("title is required")
	}
	if post.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

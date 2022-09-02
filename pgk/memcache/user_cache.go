package memcache

import (
	"template/internal/models"
)

type userModelCache []models.UserModel

var UserCache userModelCache

func (c *userModelCache) Fetch() {
	models.DB.Find(&c)
}

func (c userModelCache) Get(Username string) models.UserModel {
	for i, a := range c {
		if a.Username == Username {
			return c[i]
		}
	}
	return models.UserModel{}
}

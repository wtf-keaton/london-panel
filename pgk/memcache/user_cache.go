package memcache

import "template/internal/models"

type userModelCache []models.UserModel

var UserCache userModelCache

func (c userModelCache) Fetch() {
	models.DB.Find(&c)
}

func (c userModelCache) Get(keycode string) models.UserModel {
	for _, a := range c {
		if a.Username == keycode {
			return a
		}
	}

	return models.UserModel{}
}

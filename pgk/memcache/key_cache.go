package memcache

import "template/internal/models"

type keyModelCache []models.KeyModel

var KeyCache keyModelCache

func (c keyModelCache) Fetch() {
	models.DB.Find(&c)
}

func (c keyModelCache) Get(keycode string) models.KeyModel {
	for _, a := range c {
		if a.Keycode == keycode {
			return a
		}
	}

	return models.KeyModel{}
}

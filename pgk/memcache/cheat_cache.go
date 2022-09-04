package memcache

import "template/internal/models"

type cheatModelCache []models.CheatModel

var CheatCache cheatModelCache

func (c *cheatModelCache) Fetch() {
	models.DB.Find(&c)
}

func (c cheatModelCache) Get(Name string) models.CheatModel {
	for _, a := range c {
		if a.Name == Name {
			return a
		}
	}

	return models.CheatModel{}
}

func (c cheatModelCache) ID(Name uint) models.CheatModel {
	for _, a := range c {
		if a.ID == Name {
			return a
		}
	}

	return models.CheatModel{}
}

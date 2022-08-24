package memcache

import "template/internal/models"

type bannedHardwareModelCache []models.BannedHardware

var BannedCache bannedHardwareModelCache

func (c bannedHardwareModelCache) Fetch() {
	models.DB.Find(&c)
}

func (c bannedHardwareModelCache) Get(Name string) models.BannedHardware {
	for _, a := range c {
		if a.HardwareID == Name {
			return a
		}
	}

	return models.BannedHardware{}
}

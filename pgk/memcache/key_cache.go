package memcache

import "template/internal/models"

type keyModelCache []models.KeyModel

var KeyCache keyModelCache

func (c *keyModelCache) Fetch() {
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

func (c keyModelCache) Actived() (u int) {
	u = 0
	for _, a := range c {
		if a.Status == 1 {
			u += 1
		}
	}

	return
}

func (c keyModelCache) SellerActived(Seller string) (u int) {
	u = 0
	for _, a := range c {
		if a.CreatedBy == Seller && a.Status == 1 {
			u += 1
		}
	}

	return
}

func (c keyModelCache) SellerBanned(Seller string) (u []models.KeyModel) {
	for _, a := range c {
		if a.CreatedBy == Seller && a.Banned == true {
			u = append(u, a)
		}
	}

	return
}

func (c keyModelCache) SellerNotBanned(Seller string) (u []models.KeyModel) {
	for _, a := range c {
		if a.Banned == false && a.CreatedBy == Seller {
			u = append(u, a)
		}
	}

	return
}

func (c keyModelCache) NotBanned() (u []models.KeyModel) {
	for _, a := range c {
		if a.Banned == false {
			u = append(u, a)
		}
	}

	return
}

func (c keyModelCache) Banned() (u []models.KeyModel) {
	for _, a := range c {
		if a.Banned {
			u = append(u, a)
		}
	}

	return
}

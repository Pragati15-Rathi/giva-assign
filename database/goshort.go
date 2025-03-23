package database

func GetAllGoShorts() ([]GoShort, error) {
	var goshorts []GoShort

	tx := db.Find(&goshorts)
	if tx.Error != nil {
		return []GoShort{}, tx.Error
	}
	return goshorts, nil
}

func GetGoShort(id uint64) (GoShort, error) {
	var goshort GoShort

	tx := db.Where("id = ?", id).First(&goshort)
	if tx.Error != nil {
		return GoShort{}, tx.Error
	}

	return goshort, nil
}

func CreateGoShort(goShort GoShort) (GoShort, error) {
	tx := db.Create(&goShort)
	return goShort, tx.Error
}

func UpdateGoShort(goShort GoShort) (GoShort, error) {
	tx := db.Save(&goShort)
	return goShort, tx.Error
}

func FindByGoShortUrl(url string) (GoShort, error) {
	var goshort GoShort
	tx := db.Where("goshort = ?", url).First(&goshort)
	return goshort, tx.Error
}

func FindByRedirectUrl(url string) (GoShort, error) {
	var goshort GoShort
	tx := db.Where("redirect = ?", url).First(&goshort)
	return goshort, tx.Error
}

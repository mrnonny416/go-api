package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func GetallSavety() (models.Savety, error) {
	var Savety models.Savety
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Order("id").Find(&Savety).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Savety, err
	}

	tx.Commit()

	return Savety, nil
}

func CreateSavety(d models.Savety) (models.Savety, error) {
	tx := config.DB.Begin()
	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}



func UpdateSavety(d models.Savety) (models.Savety, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

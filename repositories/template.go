package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func GetallTemplate() ([]models.Template, error) {
	var Template []models.Template
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Order("id").Find(&Template).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Template, err
	}

	tx.Commit()

	return Template, nil
}

func GetbyidTemplate(id string) (models.Template, error) {
	var  Template models.Template
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Where("id = ?", id).Find(&Template).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Template, err
	}

	tx.Commit()

	return Template, nil
}

func CreateTemplate(d models.Template) (models.Template, error) {
	tx := config.DB.Begin()
	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}



func UpdateTemplate(d models.Template) (models.Template, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

func DeleteTemplate(id string) error {

	tx := config.DB.Begin()

	if err := tx.Debug().Where("id = ? ", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Template{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

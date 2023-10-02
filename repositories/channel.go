package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type StatusRequest struct {
	ID     *uint `json:"channelId"`
	Status *bool `json:"status"`
}

func GetallChannel() ([]models.Channel, error) {
	var Channel []models.Channel
	tx := config.DB.Begin()

	if err := tx.Debug().Preload(clause.Associations).Order("id").Limit(4).Find(&Channel).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Channel, err
	}

	tx.Commit()

	return Channel, nil
}

func Changestatus(d StatusRequest) error {
	tx := config.DB.Begin()

	if err := tx.Debug().Model(&models.Channel{}).
		Where("id = ?", d.ID).Updates(map[string]interface{}{
		"status":     d.Status,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func GetbyidChannel(id string) (models.Channel, error) {
	var  Channel models.Channel
	tx := config.DB.Begin()
	if err := tx.Debug().Preload(clause.Associations).Where("id = ?", id).Find(&Channel).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Channel, err
	}

	tx.Commit()

	return Channel, nil
}

func CreateChannel(d models.Channel) (models.Channel, error) {
	tx := config.DB.Begin()
	if err := tx.Debug().Clauses(clause.OnConflict{DoNothing: true}).Create(&d).Error; err != nil {
		tx.Commit()
		return d, err
	}

	tx.Commit()
	return d, nil
}



func UpdateChannel(d models.Channel) (models.Channel, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

// func Deleteform(id string) error {

// 	tx := config.DB.Begin()

// 	if err := tx.Debug().Where("id = ? ", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Form{}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	if err := tx.Debug().Where("form_id = ?", id).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.List_Form{}).Error; err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	tx.Commit()

// 	return nil
// }

package repositories

import (
	config "backend/config"
	"backend/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Getsaveunitreq struct {
	Datestart *string `json:"datestart"`
	Dateend   *string `json:"dateend"`
	Type      *string `json:"type"`
}

func GetallSaveunit(r Getsaveunitreq) ([]models.Saveunit, error) {
	var Saveunit []models.Saveunit
	query := `
		SELECT DATE_TRUNC('` + *r.Type + `', date) AS DATE, SUM(saveunits.unit) AS unit
		FROM saveunits
		WHERE date BETWEEN ? AND ?
		GROUP BY DATE_TRUNC('` + *r.Type + `', saveunits.date) Order by DATE`
	tx := config.DB.Begin()

	if err := tx.Debug().Raw(query, r.Datestart, r.Dateend).Find(&Saveunit).Error; err != nil {
		println(err.Error())
		tx.Commit()
		return Saveunit, err
	}

	tx.Commit()

	return Saveunit, nil
}

func CreateSaveunit(d models.Requnit) (models.Saveunit, error) {
	var Lastunit models.Saveunit
	tx := config.DB.Begin()
	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Last(&Lastunit).Error; err != nil {
		num := float32(0)
		data := models.Saveunit{
			BfUnit: &num,
			AtUnit: d.Unit,
			Unit:   d.Unit,
			Date:   d.Date,
		}
		if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error; err != nil {
			tx.Commit()
			return data, err
		}
		tx.Commit()
		return data, nil
	}
	Unit := *d.Unit - *Lastunit.AtUnit
	data := models.Saveunit{
		BfUnit: Lastunit.AtUnit,
		AtUnit: d.Unit,
		Unit:   &Unit,
		Date:   d.Date,
	}
	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error; err != nil {
		tx.Commit()
		return data, err
	}
	tx.Commit()
	return data, nil
}

func UpdateSaveunit(d models.Saveunit) (models.Saveunit, error) {
	tx := config.DB.Begin()

	if err := tx.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d).Error; err != nil {
		tx.Rollback()
		return d, err
	}

	tx.Commit()
	return d, nil
}

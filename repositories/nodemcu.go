package repositories

import (
	config "backend/config"
	"backend/models"
	"time"
)

func GetNodeMCUs() ([]models.NodeMCU, error) {
	var nodeMCUs []models.NodeMCU
	result := config.DB.Find(&nodeMCUs)
	return nodeMCUs, result.Error
}

func UpdateNodeMCU(millis string) error {
	nodeMCU := &models.NodeMCU{
		Millis:    millis,
		Timestamp: time.Now(),
	}
	result := config.DB.Model(&models.NodeMCU{}).Where("1 = 1").Updates(models.NodeMCU{Millis: nodeMCU.Millis, Timestamp: nodeMCU.Timestamp})

	return result.Error
}

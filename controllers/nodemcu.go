package controllers

import (
	"net/http"

	"backend/models"
	"backend/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveNodeMCU(router *gin.RouterGroup) {
	router.GET("/status", GetNodeMCU)
	router.PUT("/update", PutNodeMCU)
}

func GetNodeMCU(c *gin.Context) {
	nodeMCU, err := repositories.GetNodeMCUs()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "No data found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "Failed to fetch data from the database",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": nodeMCU,
	})
}

func PutNodeMCU(c *gin.Context) {
	var d models.NodeMCU
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}

	err = repositories.UpdateNodeMCU(d.Millis)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Update NodeMCU",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "NodeMCU is Successfully Updated.",
		"data": d,
	})
}

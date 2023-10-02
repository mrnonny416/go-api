package controllers

import (
	"net/http"

	"backend/models"
	repo "backend/repositories"

	"github.com/gin-gonic/gin"
)

func Savety(router *gin.RouterGroup) {
	router.GET("/all", GetallSavety)
	router.POST("/add", CreateSavety)
	router.PUT("/update", UpdateSavety)
}

func GetallSavety(c *gin.Context) {
	res, err := repo.GetallSavety()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Savety List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func CreateSavety(c *gin.Context) {
	var d models.Savety
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.CreateSavety(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Savety",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Savety is Successfully Created.",
		"data": res,
	})
}


func UpdateSavety(c *gin.Context) {
	var d models.Savety
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	res, err := repo.UpdateSavety(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Edit Savety",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Savety is Successfully Edite.",
		"data": res,
	})
}

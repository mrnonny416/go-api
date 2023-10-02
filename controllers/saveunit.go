package controllers

import (
	"net/http"

	"backend/models"
	repo "backend/repositories"

	"github.com/gin-gonic/gin"
)

func Saveunit(router *gin.RouterGroup) {
	router.POST("/all", GetallSaveunit)
	router.POST("/add", CreateSaveunit)
	router.PUT("/update", UpdateSaveunit)
}

func GetallSaveunit(c *gin.Context) {
	var d repo.Getsaveunitreq
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}

	res, err := repo.GetallSaveunit(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Saveunit List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func CreateSaveunit(c *gin.Context) {
	var d models.Requnit
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.CreateSaveunit(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Saveunit",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Saveunit is Successfully Created.",
		"data": res,
	})
}

func UpdateSaveunit(c *gin.Context) {
	var d models.Saveunit
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	res, err := repo.UpdateSaveunit(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Edit Saveunit",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Saveunit is Successfully Edite.",
		"data": res,
	})
}

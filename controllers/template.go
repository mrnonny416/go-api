package controllers

import (
	"net/http"

	"backend/models"
	repo "backend/repositories"

	"github.com/gin-gonic/gin"
)

func Template(router *gin.RouterGroup) {
	router.GET("/all", GetallTemplate)
	router.GET("/detail", GetbyidTemplate)
	router.POST("/add", CreateTemplate)
	router.PUT("/update", UpdateTemplate)
	router.DELETE("/delete", DeleteTemplate)
}

func GetallTemplate(c *gin.Context) {
	res, err := repo.GetallTemplate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Template List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func CreateTemplate(c *gin.Context) {
	var d models.Template
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.CreateTemplate(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Template",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Template is Successfully Created.",
		"data": res,
	})
}

func GetbyidTemplate(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.GetbyidTemplate(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Template",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}

func UpdateTemplate(c *gin.Context) {
	var d models.Template
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	res, err := repo.UpdateTemplate(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Edit Template",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Template is Successfully Edite.",
		"data": res,
	})
}

func DeleteTemplate(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}

	err := repo.DeleteTemplate(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Delete Temp",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 204,
		"msg":  "Temp is Successfully Delete.",
	})
}

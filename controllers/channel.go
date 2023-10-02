package controllers

import (
	"net/http"

	"backend/models"
	repo "backend/repositories"

	"github.com/gin-gonic/gin"
)

func Channels(router *gin.RouterGroup) {
	router.GET("/all", GetallChannel)
	router.GET("/detail", GetbyidChannel)
	router.POST("/add", CreateChannel)
	router.PUT("/changestatus", Changestatus)
	router.POST("/update", UpdateChannel)
}

func GetallChannel(c *gin.Context) {
	res, err := repo.GetallChannel()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed Getting Channel List",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": res,
	})
}

func Changestatus(c *gin.Context) {
	var err error
	var d repo.StatusRequest
	err = c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	err = repo.Changestatus(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Change Status",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Status is Successfully Change.",
	})
}

func CreateChannel(c *gin.Context) {
	var d models.Channel
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request",
		})
		return
	}

	res, err := repo.CreateChannel(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Created Channel",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 201,
		"msg":  "Channel is Successfully Created.",
		"data": res,
	})
}
// func Createformreturn(c *gin.Context) {
// 	var d models.Form
// 	err := c.BindJSON(&d)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"code": 400,
// 			"msg":  "Bad Request",
// 		})
// 		return
// 	}

// 	res, err := repo.Createformreturn(d)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"code": 500,
// 			"msg":  "Failed to Created Form",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"code": 201,
// 		"msg":  "Form is Successfully Created.",
// 		"data": res,
// 	})
// }

func GetbyidChannel(c *gin.Context) {
	var r = c.Query("id")
	if r == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Not Params",
		})
		return
	}
	res, err := repo.GetbyidChannel(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Getting Form",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": res,
	})
}

func UpdateChannel(c *gin.Context) {
	var d models.Channel
	err := c.BindJSON(&d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "Bad Request, Failed Getting Params",
		})
		return
	}
	res, err := repo.UpdateChannel(d)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "Failed to Edit Channel",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 201,
		"msg":  "Chanel is Successfully Edite.",
		"data": res,
	})
}

// func FormDelete(c *gin.Context) {
// 	var r = c.Query("id")
// 	if r == "" {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"code": 400,
// 			"msg":  "Bad Request, Failed Not Params",
// 		})
// 		return
// 	}

// 	err := repo.Deleteform(r)
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"code": 500,
// 			"msg":  "Failed to Delete Form",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code": 204,
// 		"msg":  "Form is Successfully Delete.",
// 	})
// }

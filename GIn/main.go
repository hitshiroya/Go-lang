package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct{
	Name string 
	Email string 
}

var data api


func main() {
	
	r := gin.Default()
	r.GET("/get",getReqHandler)
	r.POST("/post",postReqHandler)
	r.PUT("/put",putReqHandler)
	r.DELETE("/delete",deleteReqHandler)
	r.Run()
}

func getReqHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message": data,
	})
}
func postReqHandler(c *gin.Context){
	err := c.BindJSON(&data)

	if err!=nil{
		c.JSON(400,gin.H{
			"message":"data not found",
		})
	}

	c.JSON(http.StatusOK,gin.H{
		"message":data,
	})
}
func putReqHandler(c *gin.Context){
	err := c.BindJSON(&data)

	if err!=nil{
		c.JSON(400,gin.H{
			"message":"data not found",
		})
	}

	c.JSON(http.StatusOK,gin.H{
		"message":data,
	})
}

func deleteReqHandler(c *gin.Context){
	data = api{}
	c.JSON(201,gin.H{
		"message":data,
	})
}
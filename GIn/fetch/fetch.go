package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)
var url string = "https://jsonplaceholder.typicode.com/todos/1"
func getReqHandler(c *gin.Context){
	response,err := http.Get(url)

	if err!=nil{
		c.JSON(400,gin.H{
			"message":"data is not valid",
		})
	}

	defer response.Body.Close()

	body,err := io.ReadAll(response.Body)
	
	if err!=nil{
		c.JSON(400,gin.H{
			"message":"data is not valid",
		})
	}

	var dataMap map[string]interface{}

	err = json.Unmarshal(body,&dataMap)

	if err!=nil{
		c.JSON(400,gin.H{
			"message":err,
		})
	}
	
	c.JSON(200,gin.H{
		"message":dataMap,
	})
}

func main (){
	r:= gin.Default()
	r.GET("/posts",getReqHandler)
	r.Run()
}
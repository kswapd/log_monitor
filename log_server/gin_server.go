package main

import (
"encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
)
type ContainerJson struct {
        Type int
        Data struct {
        	Container_uuid string
        	Environment_id string
        	Namespace string
        	Container_name string
        	Timestamp string
        	Log_info struct {
        		Log_time string
        		Source string
        		Message string
        	}
        } 
    }





func main() {
	var containerJson ContainerJson

	myjson :=`{
	    "type": "container",	    
	    "data": {
	        "container_uuid": "container_uuid123",
	        "environment_id": "environment_id123",
	        "namespace": "namespace_123",
	        "container_name": "container_nam123",
	        "timestamp": "2016-11-15 17:00:00.134",
	        "log_info":{
	            "log_time":"2016-11-20 10:10:10.134",
	            "source":"stdout",
	            "message":"bash: vi: command not found"
	        }
	    }

	}`

	json.Unmarshal([]byte(myjson), &containerJson)
    fmt.Println(containerJson)


    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        /*c.JSON(200, gin.H{
            "message": "pong",
        })*/

        c.JSON(200, containerJson)
    })
     r.GET("/ping2", func(c *gin.Context) {

        c.String(200, "{\"message\":\"pong\"}")
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
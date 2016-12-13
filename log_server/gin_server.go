package main

import (
"encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
)
type ContainerJson struct {
        Type  string `json:"type"`
        Data struct {
        	Container_uuid string `json:"container_uuid"`
        	Environment_id string `json:"environment_id"`
        	Namespace string `json:"namespace"`
        	Container_name string `json:"container_name"`
        	Timestamp string `json:"timestamp"`
        	Log_info struct {
        		Log_time string `json:"log_time"`
        		Source string `json:"source"`
        		Message string `json:"message"`
        	} `json:"log_info"`
        }  `json:"data"`
    }




/*

{

    "query_type": "container",//or: app
    "container_uuid": "abc123",//required
    "environment_id": "def456",//required
    "start_time":"2016-12-05T08:43:00.000+00:00",//required
    "end_time":"2016-12-05T08:43:00.000+00:00",//optional,为空表示查询至最新
    "query_content":"error",//optional,为空表示查询所有
    "length_per_page":100,//required
    "page_index":1 //required
}

*/



type QueryJson struct {

    Query_type string `json:"query_type"` 
    Container_uuid string `json:"container_uuid"`
    Environment_id string `json:"environment_id"`
    Start_time string `json:"start_time"`
    End_time string `json:"end_time"`
    Query_content string `json:"query_content"`
    Length_per_page int `json:"length_per_page"`
    Page_index int `json:"page_index"`
}



func main() {
	var containerJson ContainerJson

		myjson :=`{
		  "type": "container",
		  "data": {
		    "container_uuid": "b7a37421-647e-4821-8686-aadfff162e14",
		    "environment_id": "123",
		    "namespace": "name123",
		    "container_name": "lonely_blackwell",
		    "timestamp": "2016-12-12T05:26:00.759+00:00",
		    "log_info": {
		      "log_time": "2016-12-12T05:25:57.000+00:00",
		      "source": "stdout",
		      "message": "2016/12/12 05:25:57 ### End of ScrapeGlobalStatus.\r"
		    }
		  }
		}`

type LoginCommand struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
	json.Unmarshal([]byte(myjson), &containerJson)
    fmt.Println(containerJson)


    r := gin.Default()
    r.GET("/getContainerInfo", func(c *gin.Context) {
        c.JSON(200, containerJson)
    })
    r.GET("/ping", func(c *gin.Context) {

        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
     r.GET("/ping2", func(c *gin.Context) {

        c.String(200, "{\"message\":\"pong\"}")
    })



    r.POST("/login", func(c *gin.Context) {
	    var loginCmd LoginCommand
	    c.BindJSON(&loginCmd)
	    c.JSON(200, gin.H{"user": loginCmd.Username})
	})

    r.GET("/loginget", func(c *gin.Context) {
	    var loginCmd LoginCommand
	    c.BindJSON(&loginCmd)
	    c.JSON(200, gin.H{"user": loginCmd.Username})
	})


    r.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        message := name + " is " + action
        c.String(200, message)
    })

    r.POST("/postForm", func(c *gin.Context) {

    	id := c.Query("id")
        name := c.Query("name")
        page := c.DefaultQuery("page", "0")
        message := c.PostForm("message")
        msg := "id:"+id +" page:"+ page +" name:"+name+ " message:"+message;
        c.String(200, msg);
       // name := c.PostForm("name")
        
        //fmt.Printf("id: %s; page: %s; name: %s;", id, page, name)
    })

    r.GET("/getForm", func(c *gin.Context) {

    	id := c.Query("id")
        name := c.Query("name")
        page := c.DefaultQuery("page", "0")
        msg := "id:"+id +" page:"+ page +" name:"+name;
        c.String(200, msg);
    })


    r.POST("/queryLogInfo", func(c *gin.Context) {
        var queryInfo QueryJson
	    c.BindJSON(&queryInfo)
	    //c.JSON(200, gin.H{"type": queryInfo.Query_type})
	    queryInfo.Query_type = "query is ok"


	    c.JSON(200, queryInfo)
    })


    r.POST("/queryMonitorInfo", func(c *gin.Context) {
        var queryInfo QueryJson
	    c.BindJSON(&queryInfo)
	    //c.JSON(200, gin.H{"type": queryInfo.Query_type})
	    queryInfo.Query_type = "query is ok'
	    c.JSON(200, queryInfo)
    })





    r.Run() // listen and serve on 0.0.0.0:8080
}
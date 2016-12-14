package main

import (
"encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
   // "time"
    "github.com/influxdata/influxdb/client/v2"
)

const (
    MyDB = "mydb"
    //username = "bubba"
    //password = "bumblebeetuna"
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



type QueryMonitorResultJson struct {
      Timestamp string `json:"timestamp"`
      Container_uuid string `json:"container_uuid"`
      Environment_id string `json:"environment_id"`
      Container_name string `json:"container_name"`
      Namespace string `json:"namespace"`
       Stats [] struct {
          Timestamp string `json:"timestamp"`
          Container_cpu_usage_seconds_total int `json:"container_cpu_usage_seconds_total"`
          Container_cpu_user_seconds_total int `json:"container_cpu_user_seconds_total"`
          Container_cpu_system_seconds_total int `json:"container_cpu_system_seconds_total"`
          Container_memory_usage_bytes int `json:"container_memory_usage_bytes"`
          Container_memory_limit_bytes int `json:"container_memory_limit_bytes"`
          Container_memory_cache int `json:"container_memory_cache"`
          Container_memory_rss int `json:"container_memory_rss"`
          Container_memory_swap int `json:"container_memory_swap"`
          Container_network_receive_bytes_total int `json:"container_network_receive_bytes_total"`
          Container_network_receive_packets_total int `json:"container_network_receive_packets_total"`
          Container_network_receive_packets_dropped_total int `json:"container_network_receive_packets_dropped_total"`
          Container_network_receive_errors_total int `json:"container_network_receive_errors_total"`
          Container_network_transmit_bytes_total int `json:"container_network_transmit_bytes_total"`
          Container_network_transmit_packets_total int `json:"container_network_transmit_packets_total"`
          Container_network_transmit_packets_dropped_total int `json:"container_network_transmit_packets_dropped_total"`
          Container_network_transmit_errors_total int `json:"container_network_transmit_errors_total"`
           Container_filesystem [] struct 
            {
              Container_filesystem_name string `json:"container_filesystem_name"`
              Container_filesystem_type string `json:"container_filesystem_type"`
              Container_filesystem_capacity int `json:"container_filesystem_capacity"`
              Container_filesystem_usage int `json:"container_filesystem_usage"`
            }  `json:"container_filesystem"`
          
          Container_diskio_service_bytes_async int `json:"container_diskio_service_bytes_async"`
          Container_diskio_service_bytes_read int `json:"container_diskio_service_bytes_read"`
          Container_diskio_service_bytes_sync int `json:"container_diskio_service_bytes_sync"`
          Container_diskio_service_bytes_total int `json:"container_diskio_service_bytes_total"`
          Container_diskio_service_bytes_write int `json:"container_diskio_service_bytes_write"`
          Container_tasks_state_nr_sleeping int `json:"container_tasks_state_nr_sleeping"`
          Container_tasks_state_nr_running int `json:"container_tasks_state_nr_running"`
          Container_tasks_state_nr_stopped int `json:"container_tasks_state_nr_stopped"`
          Container_tasks_state_nr_uninterruptible int `json:"container_tasks_state_nr_uninterruptible"`
          Container_tasks_state_nr_io_wait int `json:"container_tasks_state_nr_io_wait"`
        } `json:"stats"`
    }




type LoginCommand struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
// queryDB convenience function to query the database
func queryDB(cmd string)  (ret []client.Result){
c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://223.202.32.56:8086",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    q := client.NewQuery(cmd, "mydb", "ns")
    response, err := c.Query(q);
    if err == nil && response.Error() == nil {
		//fmt.Println(response.Results)
	}
    //return res, nil
    return response.Results
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


	json.Unmarshal([]byte(myjson), &containerJson)
   // fmt.Println(containerJson)


    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {

        c.JSON(200, gin.H{
            "message": "pong",
        })
    })


    r.POST("/queryLogInfo", func(c *gin.Context) {
        var queryInfo QueryJson
	    c.BindJSON(&queryInfo)
	    //c.JSON(200, gin.H{"type": queryInfo.Query_type})
	    //queryInfo.Query_type = "query is ok"
	    ret := queryDB("SELECT * FROM tx_bytes limit 2")
	    c.JSON(200, ret)
    })


    r.POST("/queryMonitorInfo", func(c *gin.Context) {
        var queryInfo QueryJson
        var monitorResult QueryMonitorResultJson

        mjson := `{
      "timestamp": "2016-12-13T06:05:55.782972713Z",
      "container_uuid": "",
      "environment_id": "1234",
      "container_name": "",
      "namespace": "",
      "stats": [
        {
          "timestamp": "2016-12-13T06:05:55.57493879Z",
          "container_cpu_usage_seconds_total": 11716012449561,
          "container_cpu_user_seconds_total": 6815830000000,
          "container_cpu_system_seconds_total": 4464280000000,
          "container_memory_usage_bytes": 44470272,
          "container_memory_limit_bytes": 18446744073709551615,
          "container_memory_cache": 32768,
          "container_memory_rss": 44437504,
          "container_memory_swap": 0,
          "container_network_receive_bytes_total": 51616316,
          "container_network_receive_packets_total": 554902,
          "container_network_receive_packets_dropped_total": 0,
          "container_network_receive_errors_total": 0,
          "container_network_transmit_bytes_total": 2330468437,
          "container_network_transmit_packets_total": 1327891,
          "container_network_transmit_packets_dropped_total": 0,
          "container_network_transmit_errors_total": 0,
          "container_filesystem": [
            {
              "container_filesystem_name": "/dev/vda1",
              "container_filesystem_type": "vfs",
              "container_filesystem_capacity": 84516978688,
              "container_filesystem_usage": 3203698688
            }
          ],
          "container_diskio_service_bytes_async": 0,
          "container_diskio_service_bytes_read": 0,
          "container_diskio_service_bytes_sync": 0,
          "container_diskio_service_bytes_total": 0,
          "container_diskio_service_bytes_write": 0,
          "container_tasks_state_nr_sleeping": 0,
          "container_tasks_state_nr_running": 0,
          "container_tasks_state_nr_stopped": 0,
          "container_tasks_state_nr_uninterruptible": 0,
          "container_tasks_state_nr_io_wait": 0
        }
      ]
    }` 

        json.Unmarshal([]byte(mjson), &monitorResult)
	    c.BindJSON(&queryInfo)
	    ret := queryDB("SELECT * FROM tx_bytes, rx_bytes limit 3; select * from /.*/ limit 1")
	    _=mjson
	    _= ret
	    fmt.Println(ret)
	    c.JSON(200, monitorResult)
    })


    r.Run() // listen and serve on 0.0.0.0:8080
}
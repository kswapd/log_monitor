package main

import (
"encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    //"time"
    "github.com/influxdata/influxdb/client/v2"
     //"reflect"
     "strconv"
    //
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

type StatsInfo struct {
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
        }



type QueryMonitorResultJson struct {
      Timestamp string `json:"timestamp"`
      Container_uuid string `json:"container_uuid"`
      Environment_id string `json:"environment_id"`
      Container_name string `json:"container_name"`
      Namespace string `json:"namespace"`
       Stats []StatsInfo  `json:"stats"`
    }




type LoginCommand struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
const (
    MyDB = "mydb"
    //username = "bubba"
    //password = "bumblebeetuna"
)


// queryDB convenience function to query the database
func queryDB(cmd string)  (ret []client.Result){
c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://223.202.32.56:8086",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    q := client.NewQuery(cmd, "db56", "ns")
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
	    ret := queryDB("SELECT * FROM tx_bytes limit 3")
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
    mystats := `

    [
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
        },
		{
          "timestamp": "2016-12-13T06:05:65.57493879Z",
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

    `


       // json.Unmarshal([]byte(mjson), &monitorResult)

    	//json.Unmarshal([]byte(mystats), &monitorResult.Stats)
    	//monitorResult.Stats = make(reflect.TypeOf(Stats) , 1)
  ret := queryDB("select * from /.*/ limit 10")


      	//monitorResult.Stats = make([]StatsInfo , 2)
    	//var info StatsInfo
    	timeStr := ""
    	timeNameStatResult := make(map[string]map[string]int)
	    c.BindJSON(&queryInfo)

      if len(ret[0].Series) > 0{
         // monitorResult.
          monitorResult.Timestamp = fmt.Sprintf("%s", ret[0].Series[0].Values[0][0])
          monitorResult.Container_uuid = fmt.Sprintf("%s", ret[0].Series[0].Values[0][14])
          monitorResult.Environment_id = fmt.Sprintf("%s", ret[0].Series[0].Values[0][24])
          monitorResult.Container_name = fmt.Sprintf("%s", ret[0].Series[0].Values[0][1])
          monitorResult.Namespace  = fmt.Sprintf("%s", ret[0].Series[0].Values[0][24])
      }
	    for index := 0; index < len(ret[0].Series); index++ {
	    	se := ret[0].Series[index]
        timeNameStatResult[se.Name] = make(map[string]int)

        for valIndex := 0; valIndex < len(se.Values); valIndex ++ {
      		timeStr = fmt.Sprintf("%s", se.Values[valIndex][0])
      		valStr := fmt.Sprintf("%s", se.Values[valIndex][28])
      		val,err := strconv.Atoi(valStr)
      		_ = err
      		fmt.Printf("%d :%s,%s,%s\n", index, se.Name, se.Values[valIndex][28], se.Values[valIndex][0])
      		//fmt.Println(reflect.TypeOf(se.Name))	
          timeNameStatResult[se.Name][timeStr] = val
        }
      }
      timeStat := make(map[string] *StatsInfo)

      for k, v := range timeNameStatResult {

         //t := v.(map[string]int)
          _ = k
         t := v
        // var info StatsInfo
         //fmt.Printf("k=%v.\n", k)
         for k1,val:= range t {
         // fmt.Printf("k1=%v, v1=%v\n", k1,val) 

          if _, ok := timeStat[k1]; !ok {
            //return true
            timeStat[k1] = new (StatsInfo)
          }
          info := timeStat[k1]
          switch k {
          case "cpu_usage_per_cpu":
            //StatsInfo.Container_cpu_usage_seconds_total = 

            break
          case "cpu_usage_system":
            info.Container_cpu_system_seconds_total = val
            break
          case "cpu_usage_total":
            info.Container_cpu_usage_seconds_total = val
            break
          case "cpu_usage_user":
            info.Container_cpu_user_seconds_total = val
            break
          case "fs_limit":

            break
          case "fs_usage":
            break
          case "load_average":
            break
          case "memory_usage":
            info.Container_memory_usage_bytes = val
            break
          case "memory_working_set":
            break
          case "rx_bytes":
            info.Container_network_receive_bytes_total = val
            break
          case "rx_errors":
            info.Container_network_receive_errors_total = val
            break
          case "tx_bytes":
            info.Container_network_transmit_bytes_total = val
            break
          case "tx_errors":
            info.Container_network_transmit_errors_total =  val
            break
          default:
            fmt.Println("Error metric name.")
          }



        }
    }

    monitorResult.Stats = make([]StatsInfo , len(timeStat))
    index := 0
    for k, _ := range timeStat{
      //fmt.Printf("%#v.\n",timeStat[k]);
      monitorResult.Stats[index] = *timeStat[k]
      index ++
    }

	
	

	    _=mjson
	    _=ret
	    _=mystats
	    _=monitorResult
	    //append(monitorResult.Stats, [])
	    //fmt.Println(len(monitorResult.Stats))
	   // fmt.Println(ret)
	    //monitorResult.Stats[0] = timeStatResult[timeStr]
	    //monitorResult.Timestamp = timeStr
	    c.JSON(200, monitorResult)
    })


    r.Run() // listen and serve on 0.0.0.0:8080
}
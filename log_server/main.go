package main

import (
	"encoding/json"
	"io"
	"net/http"
	//"github.com/bitly/go-simplejson"
	//"log"
	"fmt"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	str := `hello
world
v2.0`
	io.WriteString(w, str)
}

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


func test(){

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
	//b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	/*var f interface{}
	err := json.Unmarshal([]byte(myjson), &f)
	_ = err
	m := f.(map[string]interface{})
	 fmt.Println(m["type"])

	t := m["data"].(map[string]interface{})
	 fmt.Println(t["environment_id"])*/

	 json.Unmarshal([]byte(myjson), &containerJson)
    fmt.Println(containerJson)
    fmt.Println(containerJson.Data.Log_info.Message)

}

func main() {

	test()
	/*http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}


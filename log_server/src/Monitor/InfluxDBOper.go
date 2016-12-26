package Monitor
import (
    "github.com/influxdata/influxdb/client/v2"
    "log"
)
const (
    MyDB = "mydb"
    //username = "bubba"
    //password = "bumblebeetuna"
)


// queryDB convenience function to query the database
func QueryDB(cmd string)  (ret []client.Result){
c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://223.202.32.56:8086",
    })

    if err != nil {
        log.Fatalln("Error: ", err)
    }

    q := client.NewQuery(cmd, "db56", "rfc3339")
    response, err := c.Query(q);
    if err == nil && response.Error() == nil {
		//fmt.Println(response.Results)
	}
    //return res, nil
    return response.Results
}
package ChronosDB 
import (
    "fmt"
//    "log"
    "encoding/json"
    "strconv"
)

var prefix = "CHRONOSDBv1\t"

func Connect() {
   GetLink("127.0.0.1", 6379) 
}

func SetData(data string) {
    InputArray, err := ParseJson(data)
  
    if (err != nil) {
        return    
    }

    for _, data := range InputArray {
        hashdata := data.(map[string]interface{})
        name := ""
        dataPoints := hashdata["datapoints"]

        if hashdata["name"] == nil {
            continue    
        } else {
            name = hashdata["name"].(string) 
        }

        if dataPoints == nil {
            value, _ := (hashdata["value"].(json.Number)).Float64()

            if hashdata["timestamp"] == nil {
                //log.Fatalf("Connect failed: %s\n", err.Error()) 
                continue    
            }
            timestamp, err := (hashdata["timestamp"].(json.Number)).Int64()
            
            if err != nil {
                //log.Fatalf("Connect failed: %s\n", err.Error()) 
                continue    
            }
            keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
            SetTimeSeries(keyname, strconv.FormatFloat(value, 'f', 6, 64), offset, nil)
        } else {
            inputData := make(map[string][]interface{})
            for _, rowdata := range dataPoints.([]interface{}) {
                data := rowdata.([]interface{})
                timestamp, errT := (data[0].(json.Number)).Int64()
                value, errV := (data[1].(json.Number)).Float64()

                if errT != nil || errV != nil {
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
                    continue    
                }

                keyname, offset := generateTimeSeriesData(prefix, name , timestamp)
                inputData[keyname] = append(inputData[keyname], offset)
                inputData[keyname] = append(inputData[keyname], value)
            }
            //TODO: add a function to bulk insert                
            for k := range inputData {
                tag := []string{}
                _, err := BulkSetTimeSeries(k, inputData[k], tag)
                if err != nil {
                    //log.Fatalf("Connect failed: %s\n", err.Error()) 
                    continue    
                }
            }
            fmt.Println(inputData)
        }


    }
}


func QueryTimeSeriesData(query string)  {
}

//func AddDataPoint(timestamp unit32, data []string


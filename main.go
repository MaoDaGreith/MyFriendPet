package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/MaoDaGreith/MyFriendPet/config"
	"github.com/MaoDaGreith/MyFriendPet/repositories"
)

func main() {
	conf, err := config.NewConf()
	if err != nil {
		panic(err)
	}
	dbConn, err := repositories.GetDBConnection(conf.DBConnSettings())
	if err != nil {
		panic(err)
	}

	result := []map[string]any{}

	err = dbConn.SQL().Select("name").From("city").All(&result)
	if err != nil {
		fmt.Errorf("Unable to get results from db: %v", err)
	}
	convertedResult := convertBytesToString(result)

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Errorf("something went wrong, creating file")
	}
	fmt.Println("File has been created successfully")
	defer f.Close()

	data, err := json.MarshalIndent(convertedResult, "", "  ")
	if err != nil {
		fmt.Errorf("Failed to Marshal the data")
	}

	_, err = f.Write(data)
	if err != nil {
		fmt.Errorf("Error writing into file")
	}

	fmt.Println("Data has been written with success!!!")
}

func convertBytesToString(data []map[string]any) []map[string]any {
	result := make([]map[string]any, len(data))

	for i, v := range data {
		result[i] = make(map[string]interface{})
		for index, value := range v {
			if key, ok := value.([]byte); ok {
				result[i][index] = string(key)
			} else {
				result[i][index] = value
			}
		}
	}

	return result
}

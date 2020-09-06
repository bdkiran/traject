package api

import (
	"encoding/json"
	"net/http"

	"github.com/bdkiran/traject/utils"
)

func pageViewHandler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := parseMultipartFormData(r)
	if err != nil {
		utils.DefaultLogger.Error.Printf("Unable to parse request. Error: %s", err)
	}
	utils.DefaultLogger.Info.Println(string(jsonData))
}

//Function to parse multipart/form-data objects to maps of strings
func parseMultipartFormData(r *http.Request) ([]byte, error) {
	parseErr := r.ParseMultipartForm(0)
	if parseErr != nil {
		return nil, parseErr
	}

	//iterate though all the values
	utils.DefaultLogger.Info.Printf("-------------------------------------NEW REQUEST--------------------------------")
	mapData := make(map[string]string)
	for key, value := range r.Form {
		//utils.DefaultLogger.Info.Printf("%s : %s", key, value[0])
		mapData[key] = value[0]
	}

	newData, err := json.Marshal(mapData)
	if err != nil {
		utils.DefaultLogger.Warning.Println("An error occured when converting the map object to a json object.")
		return nil, err
	}

	return newData, nil
}

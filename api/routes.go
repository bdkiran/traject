package api

import (
	"encoding/json"
	"net/http"

	"github.com/bdkiran/traject/persist"
	"github.com/bdkiran/traject/utils"
)

//Simple get enpoint to tell used if the server is alive.
func alive(w http.ResponseWriter, r *http.Request) {
	utils.DefaultLogger.Info.Println("Alive function called.")
	const returnString = "Server is alive."
	response, _ := json.Marshal(returnString)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

//FormHandler that takes in posts request. The bodies must be url encoded data.
func formHandler(w http.ResponseWriter, r *http.Request) {
	utils.DefaultLogger.Info.Println("Form handler called.")

	//Parses form data into a json
	jsonData, err := formDataToJSONEncoded(r)
	if err != nil {
		sendResponse(w, "Invalid Request Sent.", http.StatusBadRequest)
	}
	//Json data is then stored indb
	//log.Println(string(jsonData))
	err = persist.CreateLead(jsonData)
	if err != nil {
		sendResponse(w, "Something went wrong.", http.StatusBadRequest)
		return
	}

	//Send response back
	sendResponse(w, "Valid Request Sent", http.StatusOK)
}

//Converst url-encoded form data from an http request to a json object.
func formDataToJSONEncoded(r *http.Request) ([]byte, error) {
	err := r.ParseForm()
	if err != nil {
		utils.DefaultLogger.Warning.Println("Unable to parse the form payload.")
		return nil, err
	}

	//Make a map of url encoded data
	mapData := make(map[string]string)
	for key, value := range r.Form {
		mapData[key] = value[0]
	}

	//Convert that map to a JSON object
	newData, err := json.Marshal(mapData)
	if err != nil {
		utils.DefaultLogger.Warning.Println("An error occured when converting the map object to a json object.")
		return nil, err
	}

	return newData, nil
}

//Creates and sends off a request based on the parameters passed in.
func sendResponse(w http.ResponseWriter, message string, responseCode int) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseCode)
	w.Write(response)
}

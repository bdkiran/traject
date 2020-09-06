package api

import (
	"io/ioutil"
	"net/http"

	"github.com/bdkiran/traject/persist"
	"github.com/bdkiran/traject/utils"
)

func sensorHandler(w http.ResponseWriter, r *http.Request) {
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.DefaultLogger.Error.Fatal(err)
	}
	utils.DefaultLogger.Info.Println(string(jsn))
	persist.ProduceMessage(jsn)
}

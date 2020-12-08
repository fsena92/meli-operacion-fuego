package config

import (
	"github.com/fsena92/meli-operacion-fuego/structs"
	"encoding/json"
	"io/ioutil"
	
)

/*LoadSatellites loads the satellites registered in the config json*/
func LoadSatellites(){
	var config = structs.Configuration{}
	readFile(&config)
	structs.SatellitesConfigured = config.Satellites
}


func readFile(cfg *structs.Configuration) {
	file, _ := ioutil.ReadFile("config/config.json")
	_ = json.Unmarshal([]byte(file), cfg)
} 
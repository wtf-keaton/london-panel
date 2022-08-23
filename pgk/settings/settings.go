package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type settings struct {
	DSN string
}

var Settings = &settings{
	DSN: "root:root@/cheatcompanynew?parseTime=true",
}

func Parse() {
	if configExists("config.json") == true {
		jsonFile, _ := os.Open("config.json")
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &Settings)
		return
	} else {
		file, _ := os.Create("config.json")
		defer file.Close()
		jsonData, _ := json.Marshal(Settings)
		file.WriteString(string(jsonData))
	}
}

func Export() {
	file, _ := os.Create("config.json")
	defer file.Close()
	jsonData, _ := json.Marshal(Settings)
	file.WriteString(string(jsonData))
}

func configExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

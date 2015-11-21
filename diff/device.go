package diff

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Device struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func RetrieveDevices() ([]*Device, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var devices []*Device
	err = json.NewDecoder(file).Decode(&devices)

	return devices, err
}

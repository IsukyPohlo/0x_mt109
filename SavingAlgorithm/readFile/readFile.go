package readFile

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Routes struct {
	Matrix [][]struct {
		DistanceInMeters    float32 `json:"distance_in_meters"`
		TravelTimeInMinutes float32 `json:"travel_time_in_minutes"`
	} `json:"matrix"`
}

func ReadFileJson(filePath string) (Routes, error) {

	var routesData Routes

	jsonFile, err := os.Open(filePath)

	if err != nil {
		return routesData, err
	}

	byteFile, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return routesData, err
	}

	json.Unmarshal(byteFile, &routesData)

	defer jsonFile.Close()

	return routesData, nil

}

package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func SaveToFileJSON(data any, filename string) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, fileData, 0644)
}

func CheckIfFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

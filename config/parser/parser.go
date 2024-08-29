package parser

import (
	"encoding/json"
	"os"
)

func ConfigParser(filePath string, parseStruct interface{}) error {

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, parseStruct)
	if err != nil {
		return err
	}

	return nil
}

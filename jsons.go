// Created by East.Zhang at 2022.06.03

package basic

import (
	"encoding/json"
	"os"
)

func LoadJSONFile(filename string, t any) (any, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, t)
	return t, err
}

func LoadJSONString(content string, t any) (any, error)  {
	err := json.Unmarshal([]byte(content), t)
	return t, err
}

func DumpToJSONString(t any) (string, error) {
	data, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DumpToJSONFile(filename string, t any) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0)
	return err
}
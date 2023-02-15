package config

import (
	"encoding/json"
	"os"
	"strings"
)

func LoadFile(file_name string) (config Configuration, err error) {
	var data []byte
	data, err = os.ReadFile(file_name)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	m := map[string]interface{}{}
	err = decoder.Decode(&m)
	if err == nil {
		config = &DefaultConfig{configData: m}
	}
	return
}

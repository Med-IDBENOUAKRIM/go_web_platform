package config

import "strings"

type DefaultConfig struct {
	configData map[string]interface{}
}

func (c *DefaultConfig) get(name string) (result interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		result, found = data[key]
		if newSection, ok := result.(map[string]interface{}); ok && found {
			data = newSection
		} else {
			return
		}
	}
	return
}

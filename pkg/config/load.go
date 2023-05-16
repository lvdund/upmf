package config

import (
	"encoding/json"
	"io/ioutil"
	"upmf/internal/context"
)

func LoadConfig(fn string) (cfg context.UpmfConfig, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	}
	err = json.Unmarshal(buf, &cfg)
	return
}

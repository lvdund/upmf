package config

import (
	"encoding/json"
	"io/ioutil"
	"upf/models"
)

func LoadConfig(filename string) (cfg models.UpfConfig, err error) {
	// load file config
	var buf []byte
	if buf, err = ioutil.ReadFile(filename); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &cfg)
	}

	return
}

func LoadUpmfInfo() (sbi2upmf models.Sbi, err error) {
	var buf []byte
	if buf, err = ioutil.ReadFile("config/static.json"); err != nil {
		return
	} else {
		err = json.Unmarshal(buf, &sbi2upmf)
	}
	return
}

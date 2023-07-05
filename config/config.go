package config

import (
	"encoding/json"
	"io/ioutil"
	"upf/models"
)

type UpfConfig struct {
	Id     string     `json:"id"`
	Slices []string   `json:"slices"`
	Infs   []string   `json:"infs"`
	Sbi    models.Sbi `json:"sbi"`
}

func LoadConfig(filename string) (cfg UpfConfig, err error) {
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

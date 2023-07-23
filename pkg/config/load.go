package config

import (
	"encoding/json"
	"io/ioutil"
	"upmf/internal/context"
)

func LoadConfig(fn string) (cfg *context.UpmfConfig, err error) {

	var buf []byte
	if buf, err = ioutil.ReadFile(fn); err != nil {
		return
	}
	err = json.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}

	var netnames struct {
		Networks context.NetConfig `json:"networks"`
	}
	err = json.Unmarshal(buf, &netnames)
	if err != nil {
		return
	}
	cfg.Nets = make(map[string]uint8)
	for _, name := range netnames.Networks.Access { // name: ["an1", "an2", "an3"]
		cfg.Nets[name] = context.NET_TYPE_AN
	}

	for _, name := range netnames.Networks.Transport { // ["tran"]
		cfg.Nets[name] = context.NET_TYPE_TRAN
	}

	for _, name := range netnames.Networks.Dnn { // ["e1", "e2", "internet"]
		cfg.Nets[name] = context.NET_TYPE_DNN
	}

	return
}

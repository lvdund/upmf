package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"upf/config"
	"upf/models"

	"github.com/sirupsen/logrus"
)

func RegistrationUPMF(cfg *config.UpfConfig, sbiupmf *models.Sbi) (err error) {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		logrus.Errorln("Error load config to register")
		return
	}

	// create request to upmf
	addrupmf := sbiupmf.Ip.String() + ":" + fmt.Sprint(sbiupmf.Port)
	req, err := http.NewRequest("POST", addrupmf, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorln("Fail sending request to UPMF")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Cannot register to upmf")
		return
	}
	defer resp.Body.Close()

	return nil
}

func HeartbeatRequest(sbiupmf *models.Sbi) (err error) {
	// create request to upmf
	addr2upmf := sbiupmf.Ip.String() + ":" + fmt.Sprint(sbiupmf.Port)
	req, err := http.NewRequest("POST", addr2upmf, nil)
	if err != nil {
		logrus.Errorln("Fail sending heartbeat request to UPMF")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Cannot send heartbeat request to upmf")
		return
	}
	defer resp.Body.Close()

	return nil
}

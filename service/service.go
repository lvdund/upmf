package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"upf/models"

	"github.com/sirupsen/logrus"
)

func RegistrationUPMF(cfg *models.UpfConfig, sbiupmf *models.Sbi) (err error) {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		logrus.Errorln("Error load config to register")
		return
	}
	ToFile(jsonData)

	// create request to upmf
	addr2upmf := url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(sbiupmf.Ip.String(), fmt.Sprint(sbiupmf.Port)),
		Path:   "/",
	}
	req, err := http.NewRequest(http.MethodPut, addr2upmf.String(), bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorln("Cannot setup request to UPMF:", err)
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
	logrus.Infoln("Sending register Request to UPMF:", addr2upmf.String())
	defer resp.Body.Close()

	return nil
}

func HeartbeatRequest(sbiupmf *models.Sbi) (err error) {
	// create message
	jsonString := `{"mess":"HeartbeatRequest"}`

	// create request to upmf
	addr2upmf := net.JoinHostPort(sbiupmf.Ip.String(), fmt.Sprint(sbiupmf.Port))
	req, err := http.NewRequest(http.MethodPost, addr2upmf, bytes.NewBufferString(jsonString))
	if err != nil {
		logrus.Errorln("Cannot setup heartbeat request to UPMF")
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
	logrus.Infoln("Sending heartbeat request to upmf")
	defer resp.Body.Close()

	return nil
}

func DeregistrationUPMF(id *string, sbiupmf *models.Sbi) (err error) {
	// create message
	jsonString := `{"id":"` + fmt.Sprint(id) + `"}`

	// create request to upmf
	addr2upmf := net.JoinHostPort(sbiupmf.Ip.String(), fmt.Sprint(sbiupmf.Port))
	req, err := http.NewRequest(http.MethodDelete, addr2upmf, bytes.NewBufferString(jsonString))
	if err != nil {
		logrus.Errorln("Cannot setup deregister request to UPMF")
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Cannot send deregister request to upmf")
		return
	}
	logrus.Infoln("Sending deregister request to upmf")
	defer resp.Body.Close()
	return nil
}

func ToFile(jsondata []byte) {
	file, err := os.Create("output_netinf.json")
	if err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write(jsondata)
	if err != nil {
		return
	}
}

package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"smf/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

// func QueryPath(query models.PathQuery, sbiupmf models.Sbi) (path models.DataPath, err error) {
func QueryPath() (path models.DataPath, err error) {
	// test
	pathquery := models.PathQuery{
		Dnn: "internet",
		Snssai: models.Snssai{
			Sd:  "12345",
			Sst: 0,
		},
		Nets: []string{"an1"},
	}
	query := models.Query{
		SmfId: "smf1",
		UeId: "ue1",
		Query: pathquery,
	}
	sbiupmf := models.Sbi{
		Ip:   net.ParseIP("127.0.0.1"),
		Port: 8081,
	}
	jsonBytes, _ := json.Marshal(query)

	// create request to upmf
	addr2upmf := url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(sbiupmf.Ip.String(), fmt.Sprint(sbiupmf.Port)),
		Path:   "/smf/query",
	}
	req, err := http.NewRequest(http.MethodGet, addr2upmf.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		logrus.Errorln("Cannot setup request to UPMF:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// sending request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorln("Cannot send query path to upmf:", err)
		return
	}
	logrus.Infoln("Sending query to UPMF:", addr2upmf.String())
	defer resp.Body.Close()

	// get path
	// responseData, err := ioutil.ReadAll(resp.Body)
	logrus.Infoln(resp.Body)
	err = json.NewDecoder(resp.Body).Decode(&path)
	if err != nil {
		logrus.Infoln("Error read response")
		return
	}
	logrus.Infoln("Recive:", path)
	// if err = json.Unmarshal(responseData, &path); err != nil {
	// 	logrus.Errorln("Cannot read response message from UPMF:", err)
	// 	return
	// }

	return path, err
}

func Print(path *models.DataPath) {
	spew.Config.Indent = "\t"
	str := spew.Sdump(path)
	logrus.Infof("==================================================")
	logrus.Infof("%s", str)
	logrus.Infof("==================================================")
}

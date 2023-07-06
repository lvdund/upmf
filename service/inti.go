package service

import (
	"os"
	"time"
	"upf/config"
	"upf/models"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

type UPF struct {
	Config   *models.UpfConfig `json:"config"`
	Sbi2Upmf models.Sbi        `json:"sbi2upmf"`
}

func New(cfg *models.UpfConfig) (nf *UPF, err error) {
	sbi2upmf, err := config.LoadUpmfInfo()
	if err != nil {
		return
	}
	nf = &UPF{
		Config:   cfg,
		Sbi2Upmf: sbi2upmf,
	}
	return
}

func (nf *UPF) Print() {
	spew.Config.Indent = "\t"
	str := spew.Sdump(nf)
	logrus.Infof("==================================================")
	logrus.Infof("%s", str)
	logrus.Infof("==================================================")
}

func (nf *UPF) Start() (err error) {
	logrus.Infoln("Running UPF")

	go func() {
		for {
			err = RegistrationUPMF(nf.Config, &nf.Sbi2Upmf)
			if err == nil {
				logrus.Infoln("Registered to upmf")
				break
			}
			time.Sleep(5 * time.Second)
		}
		for {
			HeartbeatRequest(&nf.Sbi2Upmf)
			time.Sleep(5 * time.Second)
		}
	}()

	return
}

func (nf *UPF) Terminate() {
	logrus.Infoln("Received a kill signal")
	DeregistrationUPMF(&nf.Config.Id, &nf.Sbi2Upmf)
	logrus.Infoln("Terminated UPF")
	os.Exit(1)
}

package service

import (
	"upf/config"
	"upf/models"
)

type UPF struct {
	Config   *config.UpfConfig `json:"config"`
	Sbi2Upmf models.Sbi        `json:"sbi2upmf"`
}

func New(cfg *config.UpfConfig) (nf *UPF, err error) {
	sbi2upmf, err := config.LoadUpmfInfo()
	if err != nil {
		return
	}
	nf = &UPF{
		Config: cfg,
		Sbi2Upmf: sbi2upmf,
	}
	return
}

func (nf *UPF) Start() (err error) {
	RegistrationUPMF(nf.Config, &nf.Sbi2Upmf)
	return
}
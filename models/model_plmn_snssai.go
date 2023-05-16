package models

type PlmnSnssai struct {
	PlmnId     *PlmnId  `json:"plmnid"`
	SNssaiList []Snssai `json:"snssailist"`
}

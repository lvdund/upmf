package models

// represent a query to search for a data path to serve a given pdu session
type PathQuery struct {
	Dnn    string   `json:"dnn"`    //to locate anchor UPF
	Snssai Snssai   `json:"snssai"` //all UPFs must serve this snssai
	Nets   []string `json:"nets"`   //names of UP networks that connect to RAN UP (to find the first RAN-connected UPF)
}
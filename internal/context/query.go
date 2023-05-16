package context

type Query struct {
	SmfId string    `json:"smfid"`
	UeId  string    `json:"ueid"`
	Query PathQuery `json:"query`
}

type QueryMap struct {
	// SmfInfo SmfNode   `json:"smfinfo"`
	// UE      UeNode    `json:"listue"`
	UeId  string    `json:"ueid"`
	Query PathQuery `json:"query`
	Paths DataPath  `json:"paths"`
}

type UeNode struct {
	Id  string    `json:"id"`
	Sbi SbiConfig `json:"sbi"`
}

type SmfNode struct {
	Id  string    `json:"id"`
	Sbi SbiConfig `json:"sbi"`
}

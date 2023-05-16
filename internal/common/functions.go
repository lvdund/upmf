package common

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"upmf/models"

	"github.com/free5gc/nas/security"
)

func BearerType(access models.AccessType) uint8 {
	if access == models.ACCESSTYPE__3_GPP_ACCESS {
		return security.Bearer3GPP
	} else if access == models.ACCESSTYPE_NON_3_GPP_ACCESS {
		return security.BearerNon3GPP
	} else {
		return security.OnlyOneBearer
	}
}

func IsSliceEqual(s1, s2 *models.Snssai) bool {
	if s1 == nil && s2 == nil {
		return true
	} else if s1 == nil || s2 == nil {
		return false
	}

	return s1.Sst == s2.Sst && strings.Compare(s1.Sd, s2.Sd) == 0
}

func IsPlmnIdEqual(id1, id2 *models.PlmnId) bool {
	if id1 == nil && id2 == nil {
		return true
	} else if id1 == nil || id2 == nil {
		return false
	}

	return strings.Compare(id1.Mnc, id2.Mnc) == 0 && strings.Compare(id1.Mcc, id2.Mcc) == 0
}
func Bytes2PlmnId(buf []byte) (id *models.PlmnId, err error) {
	if len(buf) != 3 {
		err = fmt.Errorf("plmnid must be 3-byte length")
		return
	}
	var mcc [3]byte
	var mnc [3]byte
	mcc[0] = buf[0] & 0x0f
	mcc[1] = (buf[0] & 0xf0) >> 4
	mcc[2] = (buf[1] & 0x0f)

	mnc[0] = (buf[2] & 0x0f)
	mnc[1] = (buf[2] & 0xf0) >> 4
	mnc[2] = (buf[1] & 0xf0) >> 4

	tmp := []byte{(mcc[0] << 4) | mcc[1], (mcc[2] << 4) | mnc[0], (mnc[1] << 4) | mnc[2]}

	str := hex.EncodeToString(tmp)
	plmnid := models.PlmnId{
		Mcc: str[:3],
	}
	if str[5] == 'f' {
		plmnid.Mnc = str[3:5] //discard the last letter
	} else {
		plmnid.Mnc = str[3:]
	}
	id = &plmnid
	return
}

func PlmnId2Bytes(id *models.PlmnId) (buf []uint8, err error) {
	if len(id.Mcc) != 3 {
		err = fmt.Errorf("Mcc len must be 3: %s", id.Mcc)
		return
	}
	if len(id.Mnc) != 2 && len(id.Mnc) != 3 {
		err = fmt.Errorf("Mnc len must be 2 or 3: %s", id.Mnc)
		return
	}

	var (
		mcc [3]uint8
		mnc [3]uint8
		tmp int
	)

	mnc[2] = 0x0f

	for i := 0; i < 3; i++ {
		if tmp, err = strconv.Atoi(string(id.Mcc[i])); err != nil {
			return
		}
		mcc[i] = uint8(tmp)
	}
	for i := 0; i < len(id.Mnc); i++ {
		if tmp, err = strconv.Atoi(string(id.Mnc[i])); err != nil {
			return
		}
		mnc[i] = uint8(tmp)
	}

	buf = []uint8{
		(mcc[1] << 4) | mcc[0],
		(mnc[2] << 4) | mcc[2],
		(mnc[1] << 4) | mnc[0],
	}
	return
}
func ServingNetworkName(id *models.PlmnId) string {
	//return fmt.Sprintf("5G:mnc%03x.mcc%03x.3gppnetwork.org", id.Mnc, id.Mcc)
	if len(id.Mnc) == 2 {
		return fmt.Sprintf("5G:mnc0%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	} else {
		return fmt.Sprintf("5G:mnc%s.mcc%s.3gppnetwork.org", id.Mnc, id.Mcc)
	}
}

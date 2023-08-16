package n4

import (
	// "etrib5gc/pfcp"
	// "etrib5gc/pfcp/pfcptypes"
	
	// n4msg "github.com/lvdund/n4interface"
	msgType "github.com/lvdund/n4interface/msgType"

	"smf/models"
	"math"
	"net"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "smf:up"})
}

type Upf struct {
	pfcp *pfcp.Pfcp //for sending Pfcp Requests
	ip   net.IP
	port int

	sessions map[uint64]*MsgSession

	/*
		pdrPool sync.Map
		farPool sync.Map
		barPool sync.Map
		qerPool sync.Map
		// urrPool
		// sync.Map
	*/
	pdridgen  IdGenerator
	faridgen  IdGenerator
	teididgen IdGenerator
}

func NewUpf(pfcp *pfcp.Pfcp, ip net.IP, port int) (upf *Upf) {
	upf = &Upf{
		pfcp:      pfcp,
		ip:        ip,
		port:      port,
		teididgen: NewIdGenerator(1, math.MaxUint32),
		pdridgen:  NewIdGenerator(1, math.MaxUint16),
		faridgen:  NewIdGenerator(1, math.MaxUint16),
		sessions:  make(map[uint64]*MsgSession),
	}
	return
}

func (upf *Upf) Id() string {
	return upf.ip.String()
}

func (upf *Upf) GenerateTeid() uint32 {
	return uint32(upf.teididgen.Allocate())
}

func (upf *Upf) FreeTeid(id uint32) {
	upf.teididgen.Free(uint64(id))
}
func (upf *Upf) CreateSession(localseid uint64) (session *MsgSession) {
	session = newMsgSession(localseid, upf)
	upf.sessions[localseid] = session
	return
}
func (upf *Upf) FindSession(localseid uint64) (session *MsgSession) {
	session, _ = upf.sessions[localseid]
	return
}

func (upf *Upf) createPdr() (pdr *PDR) {
	pdr = &PDR{
		PDRID: uint16(upf.pdridgen.Allocate()),
		FAR:   upf.createFar(),
	}
	return
}

func (upf *Upf) createFar() (far *FAR) {
	far = &FAR{
		FARID: uint32(upf.faridgen.Allocate()),
	}
	return
}

func (upf *Upf) removePdr(pdr *PDR) {
	upf.pdridgen.Free(uint64(pdr.PDRID))
	if pdr.FAR != nil {
		upf.removeFar(pdr.FAR)
	}
}

func (upf *Upf) removeFar(far *FAR) {
	upf.faridgen.Free(uint64(far.FARID))
}
func (upf *Upf) GetQer(authdefqos *models.AuthorizedDefaultQos) (qer *QER, err error) {
	qer = &QER{
		QERID: 1,
		QFI: msgType.QFI{
			QFI: 5,
		},
	}
	return
}

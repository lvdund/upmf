package pfcp

import (
	"upmf/internal/pfcp/pfcpmsg"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	PFCP_BUF_SIZE int = 2048
)

var log *logrus.Entry

func init() {
	log = logrus.WithFields(logrus.Fields{"mod": "pfcp"})
}

type RecvInfo struct {
	msg    *pfcpmsg.Message
	remote *net.UDPAddr
}

type Forwarder struct {
	conn *net.UDPConn
	addr net.UDPAddr
	when time.Time //started time
	wg   sync.WaitGroup
}

func newForwarder(addr net.UDPAddr) *Forwarder {
	ret := &Forwarder{
		addr: addr,
	}
	return ret
}

func (fwd *Forwarder) start(recv chan<- RecvInfo) (err error) {

	if fwd.conn, err = net.ListenUDP("udp", &fwd.addr); err != nil {
		log.Errorf("Failed to listen: %s", err.Error())
		return
	}

	log.Infof("Listen on N4 interface %s", fwd.conn.LocalAddr().String())

	go fwd.loop(recv)
	fwd.when = time.Now()
	return
}

func (fwd *Forwarder) loop(recv chan<- RecvInfo) {
	fwd.wg.Add(1)
	defer fwd.wg.Done()
	buf := make([]byte, PFCP_BUF_SIZE)
	var msg *pfcpmsg.Message
	for {
		if n, addr, err := fwd.conn.ReadFromUDP(buf); err == nil {
			msg = new(pfcpmsg.Message)
			if err = msg.Unmarshal(buf[:n]); err == nil {
				if recv != nil {
					recv <- RecvInfo{
						msg:    msg,
						remote: addr,
					}
				}
			}
		} else {
			log.Errorf(err.Error())
			break
		}
	}
	if recv != nil {
		close(recv)
	}
}
func (fwd *Forwarder) stop() {
	fwd.conn.Close()
	fwd.wg.Wait()
}

//time when the forwarder started running
func (fwd *Forwarder) When() time.Time {
	return fwd.when
}

// block util message is written to the transport or an error occurs
func (fwd *Forwarder) WriteTo(msg *pfcpmsg.Message, addr *net.UDPAddr) (err error) {
	var buf []byte
	if buf, err = msg.Marshal(); err == nil {
		_, err = fwd.conn.WriteToUDP(buf, addr)
	}
	return
}

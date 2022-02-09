package tcpservice

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/panjf2000/gnet"
)

type TcpCodec struct {
	gnet.ICodec
}

var ContinueRead = errors.New("continue read")

type DataStruct struct {
	fullLength   int
	lenNumLength int
	fullData     []byte
	NotNew       bool
	ChannelID    string
	ServerID     string
}

func (d *TcpCodec) Encode(c gnet.Conn, buf []byte) ([]byte, error) {
	buf = append(proto.EncodeVarint(uint64(len(buf))), buf...)
	return buf, nil
}

func (d *TcpCodec) Decode(c gnet.Conn) ([]byte, error) {
	r, ok := c.Context().(DataStruct)
	if !ok {
		err := c.Close()
		if err != nil {
			return nil, nil
		}
	}
	bytes := c.Read()
	if len(r.fullData) == 0 {
		var fullLength uint64
		fullLength, r.lenNumLength = proto.DecodeVarint(bytes)
		r.fullLength = int(fullLength)
		if r.fullLength == 0 {
			return nil, nil
		}
	}
	fullDataLong := len(r.fullData)
	r.fullData = append(r.fullData, bytes...)
	if len(r.fullData) >= r.fullLength+r.lenNumLength {
		c.ShiftN(r.fullLength + r.lenNumLength - fullDataLong)
		res := r.fullData[r.lenNumLength : r.fullLength+r.lenNumLength]
		r.fullData = []byte{}
		c.SetContext(r)
		return res, nil
	}
	c.ShiftN(len(bytes))
	c.SetContext(r)
	return nil, nil
}

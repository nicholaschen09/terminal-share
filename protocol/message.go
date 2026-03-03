package protocol

import (
	"encoding/binary"
	"errors"
)

const (
	MsgOutput byte = 0x01
	MsgInput  byte = 0x02
	MsgResize byte = 0x03
	MsgClose  byte = 0x04
)

var ErrTooShort = errors.New("message too short")

func Encode(msgType byte, payload []byte) []byte {
	buf := make([]byte, 1+len(payload))
	buf[0] = msgType
	copy(buf[1:], payload)
	return buf
}

func Decode(data []byte) (msgType byte, payload []byte, err error) {
	if len(data) < 1 {
		return 0, nil, ErrTooShort
	}
	return data[0], data[1:], nil
}

func EncodeResize(cols, rows uint16) []byte {
	payload := make([]byte, 4)
	binary.BigEndian.PutUint16(payload[0:2], cols)
	binary.BigEndian.PutUint16(payload[2:4], rows)
	return Encode(MsgResize, payload)
}

func DecodeResize(payload []byte) (cols, rows uint16, err error) {
	if len(payload) < 4 {
		return 0, 0, ErrTooShort
	}
	cols = binary.BigEndian.Uint16(payload[0:2])
	rows = binary.BigEndian.Uint16(payload[2:4])
	return cols, rows, nil
}

package utils

import (
	"core/utils/collections"
	"encoding/binary"
)

type ByteArray struct {
	Bytes collections.List[byte]
}

func (bt *ByteArray) Clone() ByteArray {
	clone := ByteArray{}
	clone.Bytes = append(clone.Bytes, bt.Bytes...)
	return clone
}

func (bt *ByteArray) PutByte(value byte) *ByteArray {
	bt.Bytes = append(bt.Bytes, value)
	return bt
}

func (bt *ByteArray) PutBoolean(value bool) *ByteArray {
	if value {
		bt.PutByte(1)
	} else {
		bt.PutByte(0)
	}

	return bt
}

func (bt *ByteArray) PutShort(value int) *ByteArray {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(value))
	bt.Bytes = append(bt.Bytes, bytes...)
	return bt
}

func (bt *ByteArray) PutInt(value int) *ByteArray {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))
	bt.Bytes = append(bt.Bytes, bytes...)
	return bt
}

func (bt *ByteArray) PutUint(value int64) {
	calc1 := value >> 7

	if calc1 == 0 {
		bt.PutByte(byte(((value & 127) | 128)))
		bt.PutByte(0)
		return
	}

	for value != 0 {
		bt.PutByte(byte(((value & 127) | 128)))
		value = calc1
		calc1 = calc1 >> 7
	}

	bt.PutByte(byte(value & 127))
}

func (bt *ByteArray) PutString(value string) *ByteArray {
	bt.PutShort(len(value))
	bt.Bytes = append(bt.Bytes, value...)
	return bt
}

func (bt *ByteArray) GetByte() byte {
	value := bt.Bytes[0]
	bt.Bytes = bt.Bytes[1:]
	return value
}

func (bt *ByteArray) GetBoolean() bool {
	return bt.GetByte() == 1
}

func (bt *ByteArray) GetShort() int {
	value := bt.Bytes[:2]
	bt.Bytes = bt.Bytes[2:]
	return int(binary.BigEndian.Uint16(value))
}

func (bt *ByteArray) GetInt() int {
	value := bt.Bytes[:4]
	bt.Bytes = bt.Bytes[4:]
	return int(binary.BigEndian.Uint32(value))
}

func (bt *ByteArray) GetUint() int64 {
	var value int64 = 0
	var local2 int64 = 0
	var local3 int64 = 0
	var local4 int64 = -1

	for {
		local2 = int64(bt.GetByte())
		value = (value | ((local2 & 127) << (local3 * 7)))
		local4 = (local4 << 7)
		local3++

		if !(((local2 & 128) == 128) && (local3 < 5)) {
			break
		}
	}

	if ((local4 >> 1) & value) != 0 {
		value = (value | local4)
	}

	return value
}

func (bt *ByteArray) GetString() string {
	length := bt.GetShort()
	data := bt.Bytes[:length]
	bt.Bytes = bt.Bytes[length:]
	return string(data)
}

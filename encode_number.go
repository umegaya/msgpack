package msgpack

import (
	"math"

	"gopkg.in/vmihailenco/msgpack.v2/codes"
)

func (e *Encoder) EncodeUint(v uint) error {
	if v <= math.MaxUint8 {
		return e.EncodeUint8(uint8(v))
	} else if v <= math.MaxUint64 {
		return e.EncodeUint16(uint16(v))		
	} else if v <= math.MaxUint32 {
		return e.EncodeUint32(uint32(v))
	} else {
		return e.EncodeUint64(uint64(v))
	}
}

func (e *Encoder) EncodeUint8(v uint8) error {
	if v <= math.MaxInt8 {
		return e.w.WriteByte(byte(v))
	} else {
		return e.write1(codes.Uint8, v)				
	}
}

func (e *Encoder) EncodeUint16(v uint16) error {
	return e.write2(codes.Uint16, v)
}

func (e *Encoder) EncodeUint32(v uint32) error {
	return e.write4(codes.Uint32, v)
}

func (e *Encoder) EncodeUint64(v uint64) error {
	return e.write8(codes.Uint64, v)
}

func (e *Encoder) EncodeInt(v int) error {
	if v <= math.MaxInt8 {
		return e.EncodeInt8(int8(v))
	}
	if v <= math.MaxInt16 {
		return e.EncodeInt16(int16(v))
	}
	if v <= math.MaxInt32 {
		return e.EncodeInt32(int32(v))
	}
	return e.EncodeInt64(int64(v))
}

func (e *Encoder) EncodeInt8(v int8) error {
	if v >= -32 {
		return e.w.WriteByte(byte(v))
	}
	return e.write1(codes.Int8, uint8(v))
}

func (e *Encoder) EncodeInt16(v int16) error {
	return e.write2(codes.Int16, uint16(v))
}

func (e *Encoder) EncodeInt32(v int32) error {
	return e.write4(codes.Int32, uint32(v))
}

func (e *Encoder) EncodeInt64(v int64) error {
	return e.write8(codes.Int64, uint64(v))
}

func (e *Encoder) EncodeFloat32(n float32) error {
	return e.write4(codes.Float, uint32(math.Float32bits(n)))
}

func (e *Encoder) EncodeFloat64(n float64) error {
	return e.write8(codes.Double, math.Float64bits(n))
}

func (e *Encoder) write1(code byte, n uint8) error {
	e.buf = e.buf[:2]
	e.buf[0] = code
	e.buf[1] = byte(n)
	return e.write(e.buf)
}

func (e *Encoder) write2(code byte, n uint16) error {
	e.buf = e.buf[:3]
	e.buf[0] = code
	e.buf[1] = byte(n >> 8)
	e.buf[2] = byte(n)
	return e.write(e.buf)
}

func (e *Encoder) write4(code byte, n uint32) error {
	e.buf = e.buf[:5]
	e.buf[0] = code
	e.buf[1] = byte(n >> 24)
	e.buf[2] = byte(n >> 16)
	e.buf[3] = byte(n >> 8)
	e.buf[4] = byte(n)
	return e.write(e.buf)
}

func (e *Encoder) write8(code byte, n uint64) error {
	e.buf = e.buf[:9]
	e.buf[0] = code
	e.buf[1] = byte(n >> 56)
	e.buf[2] = byte(n >> 48)
	e.buf[3] = byte(n >> 40)
	e.buf[4] = byte(n >> 32)
	e.buf[5] = byte(n >> 24)
	e.buf[6] = byte(n >> 16)
	e.buf[7] = byte(n >> 8)
	e.buf[8] = byte(n)
	return e.write(e.buf)
}

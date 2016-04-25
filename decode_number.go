package msgpack

import (
	"fmt"
	"math"
	"reflect"

	"github.com/umegaya/msgpack/codes"
)

func (d *Decoder) skipN(n int) error {
	_, err := d.readN(n)
	return err
}

func (d *Decoder) uint8() (uint8, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(c), nil
}

func (d *Decoder) uint16() (uint16, error) {
	b, err := d.readN(2)
	if err != nil {
		return 0, err
	}
	return (uint16(b[0]) << 8) | uint16(b[1]), nil
}

func (d *Decoder) uint32() (uint32, error) {
	b, err := d.readN(4)
	if err != nil {
		return 0, err
	}
	n := (uint32(b[0]) << 24) |
		(uint32(b[1]) << 16) |
		(uint32(b[2]) << 8) |
		uint32(b[3])
	return n, nil
}

func (d *Decoder) uint64() (uint64, error) {
	b, err := d.readN(8)
	if err != nil {
		return 0, err
	}
	n := (uint64(b[0]) << 56) |
		(uint64(b[1]) << 48) |
		(uint64(b[2]) << 40) |
		(uint64(b[3]) << 32) |
		(uint64(b[4]) << 24) |
		(uint64(b[5]) << 16) |
		(uint64(b[6]) << 8) |
		uint64(b[7])
	return n, nil
}


//variable length reader of uint
func (d *Decoder) uint(c byte) (uint64, error) {
	if c == codes.Nil {
		return 0, nil
	}
	if codes.IsFixedNum(c) {
		return uint64(int8(c)), nil
	}
	switch c {
	case codes.Uint8:
		n, err := d.uint8()
		return uint64(n), err
	case codes.Int8:
		n, err := d.uint8()
		return uint64(int8(n)), err
	case codes.Uint16:
		n, err := d.uint16()
		return uint64(n), err
	case codes.Int16:
		n, err := d.uint16()
		return uint64(int16(n)), err
	case codes.Uint32:
		n, err := d.uint32()
		return uint64(n), err
	case codes.Int32:
		n, err := d.uint32()
		return uint64(int32(n)), err
	case codes.Uint64, codes.Int64:
		return d.uint64()
	}
	return 0, fmt.Errorf("msgpack: invalid code %x decoding uint64", c)
}

func (d *Decoder) DecodeUint() (uint, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	n, err2 := d.uint(c)
	return uint(n), err2
}


//variable length reader of int
func (d *Decoder) int(c byte) (int64, error) {
	if c == codes.Nil {
		return 0, nil
	}
	if codes.IsFixedNum(c) {
		return int64(int8(c)), nil
	}
	switch c {
	case codes.Uint8:
		n, err := d.uint8()
		return int64(n), err
	case codes.Int8:
		n, err := d.uint8()
		return int64(int8(n)), err
	case codes.Uint16:
		n, err := d.uint16()
		return int64(n), err
	case codes.Int16:
		n, err := d.uint16()
		return int64(int16(n)), err
	case codes.Uint32:
		n, err := d.uint32()
		return int64(n), err
	case codes.Int32:
		n, err := d.uint32()
		return int64(int32(n)), err
	case codes.Uint64, codes.Int64:
		n, err := d.uint64()
		return int64(n), err
	}
	return 0, fmt.Errorf("msgpack: invalid code %x decoding int64", c)
}

func (d *Decoder) DecodeInt() (int, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	n, err2 := d.int(c)
	return int(n), err2
}


//float32
func (d *Decoder) DecodeFloat32() (float32, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return d.float32(c)
}

func (d *Decoder) float32(c byte) (float32, error) {
	if c != codes.Float {
		return 0, fmt.Errorf("msgpack: invalid code %x decoding float32", c)
	}
	b, err := d.uint32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(b), nil
}

func (d *Decoder) float32Value(value reflect.Value) error {
	v, err := d.DecodeFloat32()
	if err != nil {
		return err
	}
	value.SetFloat(float64(v))
	return nil
}

//float64
func (d *Decoder) DecodeFloat64() (float64, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return d.float64(c)
}

func (d *Decoder) float64(c byte) (float64, error) {
	if c == codes.Float {
		n, err := d.float32(c)
		return float64(n), err
	}
	if c != codes.Double {
		return 0, fmt.Errorf("msgpack: invalid code %x decoding float64", c)
	}
	b, err := d.uint64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(b), nil
}

func (d *Decoder) float64Value(value reflect.Value) error {
	v, err := d.DecodeFloat64()
	if err != nil {
		return err
	}
	value.SetFloat(v)
	return nil
}

//uint8
func (d *Decoder) DecodeUint8() (uint8, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	if codes.IsFixedNum(c) {
		return uint8(c), nil
	}
	return d.uint8()
}
func (d *Decoder) uint8Value(value reflect.Value) error {
	v, err := d.DecodeUint8()
	if err != nil {
		return err
	}
	value.SetUint(uint64(v))
	return nil
}

//uint16
func (d *Decoder) DecodeUint16() (uint16, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return d.uint16()
}
func (d *Decoder) uint16Value(value reflect.Value) error {
	v, err := d.DecodeUint16()
	if err != nil {
		return err
	}
	value.SetUint(uint64(v))
	return nil
}

//uint32
func (d *Decoder) DecodeUint32() (uint32, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return d.uint32()
}
func (d *Decoder) uint32Value(value reflect.Value) error {
	v, err := d.DecodeUint32()
	if err != nil {
		return err
	}
	value.SetUint(uint64(v))
	return nil
}

//uint64
func (d *Decoder) DecodeUint64() (uint64, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	return d.uint64()
}

func (d *Decoder) uint64Value(value reflect.Value) error {
	v, err := d.DecodeUint64()
	if err != nil {
		return err
	}
	value.SetUint(v)
	return nil
}

//uint
func (d *Decoder) uintValue(value reflect.Value) error {
	v, err := d.DecodeUint()
	if err != nil {
		return err
	}
	value.SetUint(uint64(v))
	return nil
}

//int8
func (d *Decoder) DecodeInt8() (int8, error) {
	c, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	if codes.IsFixedNum(c) {
		return int8(c), nil
	}
	n, err2 := d.uint8()
	return int8(n), err2
}
func (d *Decoder) int8Value(value reflect.Value) error {
	v, err := d.DecodeInt8()
	if err != nil {
		return err
	}
	value.SetInt(int64(v))
	return nil
}

//uint16
func (d *Decoder) DecodeInt16() (int16, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	n, err2 := d.uint64()
	return int16(n), err2
}
func (d *Decoder) int16Value(value reflect.Value) error {
	v, err := d.DecodeInt16()
	if err != nil {
		return err
	}
	value.SetInt(int64(v))
	return nil
}

//uint32
func (d *Decoder) DecodeInt32() (int32, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	n, err2 := d.uint32()
	return int32(n), err2
}
func (d *Decoder) int32Value(value reflect.Value) error {
	v, err := d.DecodeInt32()
	if err != nil {
		return err
	}
	value.SetInt(int64(v))
	return nil
}

//uint64
func (d *Decoder) DecodeInt64() (int64, error) {
	_, err := d.r.ReadByte()
	if err != nil {
		return 0, err
	}
	n, err2 := d.uint64()
	return int64(n), err2
}

func (d *Decoder) int64Value(value reflect.Value) error {
	v, err := d.DecodeInt64()
	if err != nil {
		return err
	}
	value.SetInt(v)
	return nil
}

//uint
func (d *Decoder) intValue(value reflect.Value) error {
	v, err := d.DecodeInt()
	if err != nil {
		return err
	}
	value.SetInt(int64(v))
	return nil
}

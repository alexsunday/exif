package exif

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

var (
	ErrLengthNotMatch = errors.New("rational length not match")
	ErrFormatNotMatch = errors.New("cannot match request format")
)

type Entry struct {
	Ifd        Ifd
	Tag        Tag
	Format     EntryFormat
	Components int
	Raw        []byte
	order      binary.ByteOrder
}

func (e *Entry) String() string {
	return fmt.Sprintf("<ifd:%d, tag:%d, fmt:%d, len:%d, total:%d>",
		e.Ifd, e.Tag, e.Format, e.Components, len(e.Raw))
}

func (e *Entry) ReadAsUnsignedRational() ([]UnsignedRational, error) {
	if e.Format != FormatUnsignedRational {
		return nil, ErrFormatNotMatch
	}
	var out = make([]UnsignedRational, e.Components)
	if len(e.Raw) != e.Components*8 {
		return nil, ErrLengthNotMatch
	}

	for i := 0; i != e.Components; i++ {
		out[i] = UnsignedRational{
			Numerator:   e.order.Uint32(e.Raw[i*8 : i*8+4]),
			Denominator: e.order.Uint32(e.Raw[i*8+4 : (i+1)*8]),
		}
	}

	return out, nil
}

func (e *Entry) ReadAsSignedRational() ([]SignedRational, error) {
	if e.Format != FormatSignedRational {
		return nil, ErrFormatNotMatch
	}

	var out = make([]SignedRational, e.Components)
	if len(e.Raw) != e.Components*8 {
		return nil, ErrLengthNotMatch
	}

	for i := 0; i != e.Components; i++ {
		num := e.order.Uint32(e.Raw[i*8 : i*8+4])
		den := e.order.Uint32(e.Raw[i*8+4 : (i+1)*8])
		out[i] = SignedRational{
			Numerator:   int32(num),
			Denominator: int32(den),
		}
	}

	return out, nil
}

func (e *Entry) ReadAsString() (string, error) {
	if e.Format != FormatAscii {
		return "", ErrFormatNotMatch
	}

	return string(e.Raw), nil
}

func (e *Entry) ReadAsUnsignedShort() ([]uint16, error) {
	if e.Format != FormatUnsignedShort {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != 2*e.Components {
		return nil, ErrLengthNotMatch
	}

	var out = make([]uint16, e.Components)
	for i := 0; i != e.Components; i++ {
		out[i] = uint16(e.order.Uint16(e.Raw[i*2 : (i+1)*2]))
	}

	return out, nil
}

func (e *Entry) ReadAsSignedShort() ([]int16, error) {
	if e.Format != FormatSignedShort {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != 2*e.Components {
		return nil, ErrLengthNotMatch
	}

	var out = make([]int16, e.Components)
	for i := 0; i != e.Components; i++ {
		cur := uint16(e.order.Uint16(e.Raw[i*2 : (i+1)*2]))
		out[i] = int16(cur)
	}

	return out, nil
}

func (e *Entry) GetRaw() ([]byte, error) {
	return e.Raw, nil
}

func (e *Entry) GetInt8() ([]int8, error) {
	if e.Format != FormatSignedByte {
		return nil, ErrFormatNotMatch
	}

	var out = make([]int8, len(e.Raw))
	for pos, w := range e.Raw {
		out[pos] = int8(w)
	}
	return out, nil
}

func (e *Entry) GetUint32() ([]uint32, error) {
	if e.Format != FormatUnsignedLong {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != e.Components*4 {
		return nil, ErrLengthNotMatch
	}

	var out = make([]uint32, e.Components)
	for i := 0; i != e.Components; i++ {
		cur := e.order.Uint32(e.Raw[i*4 : (i+1)*4])
		out[i] = cur
	}
	return out, nil
}

func (e *Entry) GetInt32() ([]int32, error) {
	if e.Format != FormatSignedLong {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != e.Components*4 {
		return nil, ErrLengthNotMatch
	}

	var out = make([]int32, e.Components)
	for i := 0; i != e.Components; i++ {
		cur := e.order.Uint32(e.Raw[i*4 : (i+1)*4])
		out[i] = int32(cur)
	}
	return out, nil
}

func (e *Entry) GetFloat32() ([]float32, error) {
	if e.Format != FormatFloat {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != e.Components*4 {
		return nil, ErrLengthNotMatch
	}

	var out = make([]float32, e.Components)
	for i := 0; i != e.Components; i++ {
		v := e.order.Uint32(e.Raw[i*4 : (i+1)*4])
		cur := math.Float32frombits(v)
		out[i] = cur
	}
	return out, nil
}

func (e *Entry) GetDouble64() ([]float64, error) {
	if e.Format != FormatDouble {
		return nil, ErrFormatNotMatch
	}

	if len(e.Raw) != e.Components*8 {
		return nil, ErrLengthNotMatch
	}

	var out = make([]float64, e.Components)
	for i := 0; i != e.Components; i++ {
		v := e.order.Uint64(e.Raw[i*4 : (i+1)*4])
		cur := math.Float64frombits(v)
		out[i] = cur
	}
	return out, nil
}

/*
FormatAscii => string
FormatUnsignedByte => []byte
FormatUnsignedShort => []uint16
FormatUnsignedLong => []uint32
FormatUnsignedRational []UnsignedRational
FormatSignedByte => []int8
FormatUndefined => []byte
FormatSignedShort => []int16
FormatSignedLong => []int32
FormatSignedRational => []SignedRational
FormatFloat => []float32
FormatDouble => []float64
*/
func (e *Entry) GetValue() (interface{}, error) {
	switch e.Format {
	case FormatUnsignedByte:
		return e.GetRaw()
	case FormatAscii:
		return e.ReadAsString()
	case FormatUnsignedShort:
		return e.ReadAsUnsignedShort()
	case FormatUnsignedLong:
		return e.GetUint32()
	case FormatUnsignedRational:
		return e.ReadAsUnsignedRational()
	case FormatSignedByte:
		return e.GetInt8()
	case FormatUndefined:
		return e.GetRaw()
	case FormatSignedShort:
		return e.ReadAsSignedShort()
	case FormatSignedLong:
		return e.GetInt32()
	case FormatSignedRational:
		return e.ReadAsSignedRational()
	case FormatFloat:
		return e.GetFloat32()
	case FormatDouble:
		return e.GetDouble64()
	}

	return nil, ErrUnknownFormat
}

// ifd: [0: 2], tag: [2: 4], littleEndian
type IfdTag [4]byte

func (m *IfdTag) String() string {
	return fmt.Sprintf("<ifd: %d, tag: %d>", m.Ifd(), m.Tag())
}

func NewIfdTag(ifd, tag uint16) IfdTag {
	var out [4]byte
	binary.LittleEndian.PutUint16(out[:2], ifd)
	binary.LittleEndian.PutUint16(out[2:4], tag)

	return out
}

func (m *IfdTag) Ifd() uint16 {
	return binary.LittleEndian.Uint16(m[:2])
}

func (m *IfdTag) Tag() uint16 {
	return binary.LittleEndian.Uint16(m[2:4])
}

type SignedRational struct {
	Numerator   int32
	Denominator int32
}

type UnsignedRational struct {
	Numerator   uint32
	Denominator uint32
}

func (u *UnsignedRational) String() string {
	return fmt.Sprintf("%d/%d", u.Numerator, u.Denominator)
}

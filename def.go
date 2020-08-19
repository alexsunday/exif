package exif

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrLengthNotMatch = errors.New("rational length not match")
)

type EntryFormat int
type Tag uint16
type Ifd uint16

type Entry struct {
	Ifd Ifd
	Tag Tag
	Format EntryFormat
	Components int
	Raw []byte
}

func (e *Entry) String() string {
	return fmt.Sprintf("<ifd:%d, tag:%d, fmt:%d, len:%d, total:%d>",
		e.Ifd, e.Tag, e.Format, e.Components, len(e.Raw))
}

func (e *Entry) UnsignedRational(order binary.ByteOrder) ([]UnsignedRational, error) {
	var out = make([]UnsignedRational, e.Components)
	if len(e.Raw) != e.Components * 8 {
		return nil, ErrLengthNotMatch
	}

	for i:=0; i!=e.Components; i++ {
		out[i] = UnsignedRational{
			Numerator: order.Uint32(e.Raw[ i * 8: i * 8 + 4]),
			Denominator: order.Uint32(e.Raw[i * 8 + 4: (i + 1) * 8]),
		}
	}

	return out, nil
}

func (e *Entry) SignedRational(order binary.ByteOrder) ([]SignedRational, error) {
	var out = make([]SignedRational, e.Components)
	if len(e.Raw) != e.Components * 8 {
		return nil, ErrLengthNotMatch
	}

	for i:=0; i!=e.Components; i++ {
		num := order.Uint32(e.Raw[ i * 8: i * 8 + 4])
		den := order.Uint32(e.Raw[i * 8 + 4: (i + 1) * 8])
		out[i] = SignedRational{
			Numerator: int32(num),
			Denominator: int32(den),
		}
	}

	return out, nil
}

// ifd: [0: 2], tag: [2: 4], littleEndian
type IfdTag [4]byte

func (m *IfdTag) String() string {
	return fmt.Sprintf("<ifd: %d, tag: %d>", m.Ifd(), m.Tag())
}

func NewIfdTag(ifd, tag uint16) IfdTag {
	var out [4]byte
	binary.LittleEndian.PutUint16(out[:2], ifd)
	binary.LittleEndian.PutUint16(out[2: 4], tag)

	return out
}

func (m *IfdTag) Ifd() uint16 {
	return binary.LittleEndian.Uint16(m[:2])
}

func (m *IfdTag) Tag() uint16 {
	return binary.LittleEndian.Uint16(m[2: 4])
}

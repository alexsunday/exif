package exif

import (
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrLengthNotMatch = errors.New("rational length not match")
	ErrFormatNotMatch = errors.New("cannot match request format")
)

type EntryFormat int
type Tag uint16
type Ifd uint16

const (
	FormatByte EntryFormat = 1
	FormatAscii EntryFormat = 2
	FormatShort EntryFormat = 3
	FormatLong EntryFormat = 4
	FormatRational EntryFormat = 5
	FormatSignedByte EntryFormat = 6
	FormatUndefined EntryFormat = 7
	FormatSignedShort EntryFormat = 8
	FormatSignedLong EntryFormat = 9
	FormatSignedRational EntryFormat = 10
	FormatFloat EntryFormat = 11
	FormatDouble EntryFormat = 12
)

type Entry struct {
	Ifd Ifd
	Tag Tag
	Format EntryFormat
	Components int
	Raw []byte
	order binary.ByteOrder
}

func (e *Entry) String() string {
	return fmt.Sprintf("<ifd:%d, tag:%d, fmt:%d, len:%d, total:%d>",
		e.Ifd, e.Tag, e.Format, e.Components, len(e.Raw))
}

func (e *Entry) UnsignedRational() ([]UnsignedRational, error) {
	if e.Format != FormatRational {
		return nil, ErrFormatNotMatch
	}
	var out = make([]UnsignedRational, e.Components)
	if len(e.Raw) != e.Components * 8 {
		return nil, ErrLengthNotMatch
	}

	for i:=0; i!=e.Components; i++ {
		out[i] = UnsignedRational{
			Numerator: e.order.Uint32(e.Raw[ i * 8: i * 8 + 4]),
			Denominator: e.order.Uint32(e.Raw[i * 8 + 4: (i + 1) * 8]),
		}
	}

	return out, nil
}

func (e *Entry) SignedRational() ([]SignedRational, error) {
	if e.Format != FormatSignedRational {
		return nil, ErrFormatNotMatch
	}

	var out = make([]SignedRational, e.Components)
	if len(e.Raw) != e.Components * 8 {
		return nil, ErrLengthNotMatch
	}

	for i:=0; i!=e.Components; i++ {
		num := e.order.Uint32(e.Raw[ i * 8: i * 8 + 4])
		den := e.order.Uint32(e.Raw[i * 8 + 4: (i + 1) * 8])
		out[i] = SignedRational{
			Numerator: int32(num),
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

func (e *Entry) ReadAsUnsignedShort() (uint16, error) {
	if e.Format != FormatShort {
		return 0, ErrFormatNotMatch
	}

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

type SignedRational struct {
	Numerator int32
	Denominator int32
}

type UnsignedRational struct {
	Numerator uint32
	Denominator uint32
}

func (u *UnsignedRational) String() string {
	return fmt.Sprintf("%d/%d", u.Numerator, u.Denominator)
}

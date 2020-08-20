package exif

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrNotFoundEntry = errors.New("cannot found special entry")
	ErrUnknownFormat = errors.New("unknown entry data format")
	ErrValueNotMatch = errors.New("read value format not match")
	ErrValueTooSmall = errors.New("value length too small")
)

type Helper struct {
	Raw   map[IfdTag]Entry
	Order binary.ByteOrder
}

func NewHelper(data *Data) *Helper {
	return &Helper{
		Raw:   data.Raw,
		Order: data.Order,
	}
}

func (h *Helper) GetLatitude() (float64, error) {
	v, err := h.GetValue(IfdGps, EXIF_TAG_GPS_LATITUDE)
	if err != nil {
		return 0, err
	}

	rs, ok := v.([]UnsignedRational)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(rs) != 3 {
		return 0, ErrLengthNotMatch
	}

	deg := float64(rs[0].Numerator) / float64(rs[0].Denominator)
	min := float64(rs[1].Numerator) / float64(rs[1].Denominator)
	sec := float64(rs[2].Numerator) / float64(rs[2].Denominator)
	out := deg + min/60.0 + sec/3600.0

	refV, err := h.GetValue(IfdGps, EXIF_TAG_GPS_LATITUDE_REF)
	if err != nil {
		return 0, err
	}
	ref, ok := refV.(string)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(ref) == 0 {
		return 0, ErrValueTooSmall
	}
	if strings.ToUpper(ref[:1])[0] == 'S' {
		out = -1 * out
	}

	return out, nil
}

func (h *Helper) GetLongitude() (float64, error) {
	v, err := h.GetValue(IfdGps, EXIF_TAG_GPS_LONGITUDE)
	if err != nil {
		return 0, err
	}

	rs, ok := v.([]UnsignedRational)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(rs) != 3 {
		return 0, ErrLengthNotMatch
	}

	deg := float64(rs[0].Numerator) / float64(rs[0].Denominator)
	min := float64(rs[1].Numerator) / float64(rs[1].Denominator)
	sec := float64(rs[2].Numerator) / float64(rs[2].Denominator)
	out := deg + min/60.0 + sec/3600.0

	refV, err := h.GetValue(IfdGps, EXIF_TAG_GPS_LONGITUDE_REF)
	if err != nil {
		return 0, err
	}
	ref, ok := refV.(string)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(ref) == 0 {
		return 0, ErrValueTooSmall
	}
	if strings.ToUpper(ref[:1])[0] == 'W' {
		out = -1 * out
	}

	return out, nil
}

func (h *Helper) GetAltitude() (float64, error) {
	v, err := h.GetValue(IfdGps, EXIF_TAG_GPS_ALTITUDE)
	if err != nil {
		return 0, err
	}

	rs, ok := v.([]UnsignedRational)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(rs) == 0 {
		return 0, ErrLengthNotMatch
	}
	r := rs[0]

	val := float64(r.Numerator) / float64(r.Denominator)

	refV, err := h.GetValue(IfdGps, EXIF_TAG_GPS_ALTITUDE_REF)
	if err != nil {
		return 0, err
	}
	ref, ok := refV.([]byte)
	if !ok {
		return 0, ErrValueNotMatch
	}
	if len(ref) == 0 {
		return 0, ErrValueTooSmall
	}
	if ref[0] == 1 {
		val = -1 * val
	}

	return val, nil
}

type Location struct {
	Longitude float64
	Latitude  float64
	Altitude  float64
}

func (l *Location) String() string {
	return fmt.Sprintf("<Lat: %.8f, Lng: %.8f, Alt: %.4f>",
		l.Latitude, l.Longitude, l.Altitude,
	)
}

func (h *Helper) GetLocation() (*Location, error) {
	lat, err := h.GetLatitude()
	if err != nil {
		return nil, err
	}

	lng, err := h.GetLongitude()
	if err != nil {
		return nil, err
	}

	alt, err := h.GetAltitude()
	if err != nil {
		return nil, err
	}

	return &Location{
		Latitude:  lat,
		Longitude: lng,
		Altitude:  alt,
	}, nil
}

func (h *Helper) GetTimestamp() (*time.Time, error) {
	return nil, nil
}

func (h *Helper) GetEntry(ifd, tag uint16) *Entry {
	key := NewIfdTag(ifd, tag)
	v, ok := h.Raw[key]
	if ok {
		return &v
	}

	return nil
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
func (h *Helper) GetValue(ifd Ifd, tag Tag) (interface{}, error) {
	entry := h.GetEntry(uint16(ifd), uint16(tag))
	if entry == nil {
		return nil, ErrNotFoundEntry
	}

	return entry.GetValue()
}

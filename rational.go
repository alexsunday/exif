package exif

import "fmt"

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

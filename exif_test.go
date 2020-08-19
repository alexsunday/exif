package exif

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	exif := New()

	// http://www.exif.org/samples/fujifilm-mx1700.jpg
	err := exif.Open("_examples/resources/testlocation.jpg")
	assert.NoError(t, err)
	assert.True(t, len(exif.Tags) > 0)

	for key, val := range exif.Tags {
		fmt.Printf("%s: %s\n", key, val)
	}

	for key, val := range exif.Raw {
		fmt.Printf("%s: %d,%d\n", key.String(), val.Components, len(val.Raw))
	}
}

func TestRead(t *testing.T) {
	// http://www.exif.org/samples/fujifilm-mx1700.jpg
	exif, err := Read("_examples/resources/test.jpg")

	assert.NoError(t, err)
	assert.True(t, len(exif.Tags) > 0)

	for key, val := range exif.Tags {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func TestWriteAndParse(t *testing.T) {
	exif := New()

	// http://www.exif.org/samples/fujifilm-mx1700.jpg
	file, err := os.Open("_examples/resources/test.jpg")

	assert.NoError(t, err)

	defer file.Close()

	_, err = io.Copy(exif, file)

	assert.Error(t, err)
	assert.Equal(t, ErrFoundExifInData, err)

	err = exif.Parse()
	assert.NoError(t, err)

	for key, val := range exif.Tags {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func TestGetLongitude(t *testing.T) {
	exif := New()
	err := exif.Open("_examples/resources/testlocation.jpg")
	assert.NoError(t, err)

	longitude, ok := exif.Tags["Longitude"]
	assert.True(t, ok)

	assert.Equal(t, "131,  0, 55.2063", longitude)
}

func TestGetLatitude(t *testing.T) {
	e := New()
	err := e.Open("_examples/resources/testlocation.jpg")
	assert.NoError(t, err)
	latitude, ok := e.Tags["Latitude"]
	assert.True(t, ok)

	assert.Equal(t, "25, 21, 32.6101", latitude)

	entry := e.GetEntry(3, 0x02)
	assert.NotNil(t, entry)

	rs, err := entry.UnsignedRational(e.Order)
	assert.Nil(t, err)
	assert.Equal(t, len(rs), 3)
	assert.Equal(t, rs[0].Numerator, uint32(25))
	assert.Equal(t, rs[1].Numerator, uint32(21))
	assert.Equal(t, rs[2].Numerator, uint32(326101))
	assert.Equal(t, rs[2].Denominator, uint32(10000))
}

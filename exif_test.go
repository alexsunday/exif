package exif

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	exif := New()

	// http://www.exif.org/samples/fujifilm-mx1700.jpg
	f1 := "e:/PHOTO/照相馆/20170604/修/IMG_9292.jpg"
	//f1 := "_examples/resources/testlocation.jpg"
	err := exif.Open(f1)
	assert.NoError(t, err)

	for key, val := range exif.Raw {
		fmt.Printf("%s: %d,%d\n", key.String(), val.Components, len(val.Raw))
	}
}

func TestReadRaw(t *testing.T) {
	f1 := "D:\\alexs\\Documents\\IMG_20121230_140323.jpg"
	f, err := os.Open(f1)
	require.Nil(t, err)

	e := New()
	_, err = io.Copy(e, f)
	require.Equal(t, err, ErrFoundExifInData)

	err = e.Parse()
	require.Nil(t, err)

	for key, val := range e.Raw {
		fmt.Printf("%s => %s\n", key.String(), val.String())
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

	for key, val := range exif.Raw {
		fmt.Printf("%s: %s\n", key, val.String())
	}
}

func TestReadLocation(t *testing.T) {
	exif := New()
	f1 := "_examples/resources/testlocation.jpg"
	err := exif.Open(f1)
	require.Nil(t, err)

	helper := NewHelper(exif)
	loc, err := helper.GetLocation()
	require.Nil(t, err)
	println(loc.String())
}
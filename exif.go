// Copyright (c) 2012-2015 José Carlos Nieto, https://menteslibres.net/xiam
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Package exif provides bindings for libexif.
package exif

/*
#include <stdlib.h>
#include <libexif/exif-data.h>
#include <libexif/exif-loader.h>
#include <libexif/exif-content.h>
#include <libexif/exif-byte-order.h>
*/
import "C"

import (
	"encoding/binary"
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

// Error messages.
var (
	ErrNoExifData      = errors.New(`no exif data found`)
	ErrFoundExifInData = errors.New(`found exif header. OK to call Parse`)
)

// Data stores the EXIF tags of a file.
type Data struct {
	exifLoader *C.ExifLoader
	Raw        map[IfdTag]Entry
	Order      binary.ByteOrder
}

// New creates and returns a new exif.Data object.
func New() *Data {
	data := &Data{
		Raw: make(map[IfdTag]Entry),
	}
	return data
}

// Read attempts to read EXIF data from a file.
func Read(file string) (*Data, error) {
	data := New()
	if err := data.Open(file); err != nil {
		return nil, err
	}
	return data, nil
}

// Open opens a file path and loads its EXIF data.
func (d *Data) Open(file string) error {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))

	exifData := C.exif_data_new_from_file(cfile)

	if exifData == nil {
		return ErrNoExifData
	}
	defer C.exif_data_unref(exifData)

	return d.parseRaw(exifData)
}

func (d *Data) parseRaw(ed *C.ExifData) error {
	var raw []byte
	var tag uint16 = 0
	var ifd uint16 = 0

	order := C.exif_data_get_byte_order(ed)
	if order == C.EXIF_BYTE_ORDER_MOTOROLA {
		d.Order = binary.BigEndian
	} else if order == C.EXIF_BYTE_ORDER_INTEL {
		d.Order = binary.LittleEndian
	}

	for i:=0; i!= C.EXIF_IFD_COUNT; i++ {
		content := (*ed).ifd[i]
		length := int((*content).count)
		var pEntries **C.ExifEntry = (*content).entries

		if pEntries == nil {
			continue
		}
		sEntries := (*[1<<30] *C.ExifEntry)(unsafe.Pointer(pEntries))[:length:length]
		for _, pEntry := range sEntries {
			entry := *pEntry
			tag = uint16(C.uint16_t(entry.tag))

			if pEntry == nil {
				ifd = uint16(C.uint16_t(C.EXIF_IFD_COUNT))
			} else {
				ifd = uint16(C.uint16_t(C.exif_content_get_ifd(entry.parent)))
			}
			key := NewIfdTag(ifd, tag)

			if entry.data != nil && entry.size != 0 {
				raw = C.GoBytes(unsafe.Pointer(entry.data), C.int(entry.size))
			}

			d.Raw[key] = Entry{
				Ifd: Ifd(ifd),
				Tag: Tag(tag),
				Format: EntryFormat(int(C.int(entry.format))),
				Components: int(C.ulong(entry.components)),
				Raw: raw,
				order: d.Order,
			}
		}
	}

	return nil
}

// Write writes bytes to the exif loader. Sends ErrFoundExifInData error when
// enough bytes have been sent.
func (d *Data) Write(p []byte) (n int, err error) {
	if d.exifLoader == nil {
		d.exifLoader = C.exif_loader_new()
		runtime.SetFinalizer(d, (*Data).cleanup)
	}

	res := C.exif_loader_write(d.exifLoader, (*C.uchar)(unsafe.Pointer(&p[0])), C.uint(len(p)))

	if res == 1 {
		return len(p), nil
	}
	return len(p), ErrFoundExifInData
}

// Parse finalizes the data loader and sets the tags
func (d *Data) Parse() error {
	defer d.cleanup()

	exifData := C.exif_loader_get_data(d.exifLoader)
	if exifData == nil {
		return fmt.Errorf(ErrNoExifData.Error(), "")
	}

	defer func() {
		C.exif_data_unref(exifData)
	}()

	return d.parseRaw(exifData)
}

func (d *Data) cleanup() {
	if d.exifLoader != nil {
		C.exif_loader_unref(d.exifLoader)
		d.exifLoader = nil
	}
}

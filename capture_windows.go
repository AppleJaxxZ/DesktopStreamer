//go:build windows
// +build windows

package main

import (
	"bytes"
	"image/jpeg"
	"log"

	"github.com/kbinani/screenshot"
)

func CaptureScreenJPEG() []byte {
	n := screenshot.NumActiveDisplays()
	if n == 0 {
		return nil
	}

	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Println("capture error:", err)
		return nil
	}

	var buf bytes.Buffer
	jpeg.Encode(&buf, img, &jpeg.Options{Quality: 50})
	return buf.Bytes()
}

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	seed := "GnomeChild"

	width := 6
	height := 9

	img := drawSpridenticon(width, height, seed)

	// Resize image to readable resolution
	newImage := resize.Resize(160, 0, img, resize.NearestNeighbor)

	// Print image to golang play console
	displayImage(newImage)

	f, _ := os.Create(seed + ".png")
	png.Encode(f, newImage)
}

// displayImage : renders an image to the playground's console by
// base64-encoding the encoded image and printing it to stdout
// with the prefix "IMAGE:".
func displayImage(m image.Image) {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	fmt.Println("IMAGE:" + base64.StdEncoding.EncodeToString(buf.Bytes()))
}

// drawSpridenticon : draws a sprite of given width and height based off of provided seed
func drawSpridenticon(width int, height int, seed string) image.Image {
	// Using hashing algorithm to convert string to 32 bits, to be used as a seed
	crc32InUint32 := crc32.ChecksumIEEE([]byte(seed))
	crc32InInt64 := int64(crc32InUint32)

	rand.Seed(crc32InInt64)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	r := uint8(rand.Intn(100) + 100)
	g := uint8(rand.Intn(100) + 100)
	b := uint8(rand.Intn(100) + 100)

	randomColour := color.RGBA{r, g, b, 0xff}

	mid := int(int64(width-1) / int64(2))

	for x := 0; x <= mid; x++ {
		for y := 0; y < height; y++ {
			percent := rand.Intn(100)
			switch {
			case percent > 50:
				// Draw pixels on one side of the canvas
				img.Set(x, y, randomColour)
				// Mirror the pixels horizontally for symmetry
				img.Set(width-x-1, y, randomColour)
			default:
				// Use zero value.
			}
		}
	}

	return img
}

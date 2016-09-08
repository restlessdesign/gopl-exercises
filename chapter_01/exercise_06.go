package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var colorPalette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	color.RGBA{0x00, 0x00, 0x10, 0xFF},
	color.RGBA{0x00, 0x00, 0x20, 0xFF},
	color.RGBA{0x00, 0x00, 0x30, 0xFF},
	color.RGBA{0x00, 0x00, 0x40, 0xFF},
	color.RGBA{0x00, 0x00, 0x50, 0xFF},
	color.RGBA{0x00, 0x00, 0x60, 0xFF},
	color.RGBA{0x00, 0x00, 0x70, 0xFF},
	color.RGBA{0x00, 0x00, 0x80, 0xFF},
	color.RGBA{0x00, 0x00, 0x90, 0xFF},
	color.RGBA{0x00, 0x00, 0xA0, 0xFF},
	color.RGBA{0x00, 0x00, 0xB0, 0xFF},
	color.RGBA{0x00, 0x00, 0xC0, 0xFF},
	color.RGBA{0x00, 0x00, 0xD0, 0xFF},
	color.RGBA{0x00, 0x00, 0xE0, 0xFF},
}

var colorIndex = rand.Intn(255)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// Like vars, consts can be defined inside or outside of functions
	const (
		cycles  = 3     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size...+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames
	)

	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // gif.GIF is a struct
	phase := 0.0                        // phase difference
	phaseInc := 0.1                     // phase incrementation

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, colorPalette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			srand := rand.New(rand.NewSource(int64(y * 10)))
			randIndex := srand.Intn(len(colorPalette))

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randIndex))
		}

		phase += phaseInc

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Note: ignoring encoding errors here
	gif.EncodeAll(out, &anim)
}

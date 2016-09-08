package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/server-info", serverInfoHandler)
	http.HandleFunc("/gif", gifHandler)
	http.HandleFunc("/favicon.ico", gifHandler)

	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

var count int
var mu sync.Mutex

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count: %d", count)
	mu.Unlock()
}

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "Remote Address: %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func gifHandler(w http.ResponseWriter, r *http.Request) {
	var ncycles int = 5

	if err := r.ParseForm(); err == nil {
		cycles, err := strconv.Atoi(r.FormValue("cycles"))

		if err == nil {
			ncycles = cycles
		}
	}

	lissajous(w, ncycles)
}

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xff, 0x00, 0xff},
}

// Note: Constant values may only be numbers, strings, or booleans!
const (
	backgroundIndex = 0
	foregroundIndex = 1
)

type CycleGenerator struct {
	cycles int
}

func lissajous(out io.Writer, cycles int) {
	// Like vars, consts can be defined inside or outside of functions
	const (
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
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), foregroundIndex)
		}

		phase += phaseInc

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Note: ignoring encoding errors here
	gif.EncodeAll(out, &anim)
}

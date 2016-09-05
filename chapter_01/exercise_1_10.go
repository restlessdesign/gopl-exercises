package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	f, err := os.OpenFile("/tmp/fetch_results", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file %s: %v", f, err)
		return
	}
	defer f.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// Receive from channel
	for range os.Args[1:] {
		val := <-ch
		fmt.Println(val)
		f.WriteString(val + "\n")
	}

	fmt.Printf("%.3f seconds elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	if strings.HasPrefix(url, "http://") != true {
		url = "http://" + url
	}

	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// Send response to /dev/null (we’re just interested in timings here)
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error reading %s: %v", url, err)
	}

	duration := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%d  %s      %7d %.3f", resp.StatusCode, url, nbytes, duration)
}

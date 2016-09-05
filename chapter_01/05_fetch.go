package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
			os.Exit(1)
		}

		bod, err := ioutil.ReadAll(resp.Body)

		// Be sure to close body
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: Reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", bod)
	}
}

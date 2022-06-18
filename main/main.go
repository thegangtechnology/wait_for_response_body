package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var url = flag.String("url", "http://localhost/", "URL to poll")
	var responseCode = flag.Int("code", 200, "Response code to wait for")
	var timeout = flag.Int("timeout", 2000, "Timeout before giving up in ms")
	var interval = flag.Int("interval", 200, "Interval between polling in ms")
	var expectedResponse = flag.String("expectedResponse", "", "Expected response of the endpoint")
	var localhost = flag.String("localhost", "", "Ip address to use for localhost")
	flag.Parse()

	fmt.Printf("Polling URL `%s` for response code %d for up to %d ms at %d ms intervals\n", *url, *responseCode, *timeout, *interval)
	startTime := time.Now()
	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	sleepDuration := time.Duration(*interval) * time.Millisecond

	if *localhost!="" && strings.Contains(*url, "localhost") {
		*url = strings.ReplaceAll(*url, "localhost", *localhost)
	}
	for {
		res, err := http.Get(*url)

		if err == nil && res.StatusCode == *responseCode {
			// check response body
			if expectedResponse != "" {
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if body == expectedResponse {
					fmt.Printf("Response body: %v", body)
					os.Exit(0)
				}
			} else {
				fmt.Printf("Response header: %v", res)
				os.Exit(0)
			}
		}

		time.Sleep(sleepDuration)
		elapsed := time.Now().Sub(startTime)

		if elapsed > timeoutDuration {
			fmt.Printf("Timed out\n")
			os.Exit(1)
		}
	}
}

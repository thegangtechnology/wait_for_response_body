package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	var interval = flag.Int("interval", 200, "Interval between polling in ms")
	var method = flag.String("method", "HEAD", "HTTP request method")
	var pollUrl = flag.String("url", "http://localhost/", "URL to poll")
	var expectedBody = flag.String("expectedbody", "", "Expected response of the endpoint")
	var expectedCode = flag.Int("expectedcode", 200, "Response code to wait for")
	var server = flag.String("server", "", "Ip address/hostname to connect to")
	var timeout = flag.Int("timeout", 2000, "Timeout before giving up in ms")
	flag.Parse()

	u, err := url.Parse(*pollUrl)
	if err != nil {
		logrus.Fatalf("fatal error while parsing the url: %s", *pollUrl)
	}
	// Save the hostname to use as the HTTP Host header before pointing at the server, if present
	httpHost := u.Hostname()
	if server != nil {
		u.Host = fmt.Sprintf("%s:%s", *server, u.Port())
	}
	logrus.Infof("Polling URL `%s` with %s method for response code %d against %s for up to %d ms at %d ms intervals", u, *method, *expectedCode, *server, *timeout, *interval)
	logrus.Debugf("HTTP Host: %s", httpHost)

	startTime := time.Now()
	timeoutDuration := time.Duration(*timeout) * time.Millisecond
	sleepDuration := time.Duration(*interval) * time.Millisecond

	for {
		req, err := http.NewRequest(*method, u.String(), nil)
		if err != nil {
			logrus.Errorf("failed to set up the http connection: %s", err)
			os.Exit(1)
		}
		req.Host = httpHost
		logrus.Debugf("Requesting...")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			logrus.Errorf("client: error making http request: %s", err)
			time.Sleep(sleepDuration)
			continue
		}

		logrus.Debugf("Response code - Want: %d - Got: %d", *expectedCode, res.StatusCode)
		if res.StatusCode != *expectedCode {
			logrus.Errorf("Invalid response code received: expected %d, got %d", *expectedCode, res.StatusCode)
			os.Exit(1)
		}
		b, _ := io.ReadAll(res.Body)
		logrus.Debugf("body: %v", strings.Trim(string(b), "\n"))
		if *expectedBody != "" {
			defer res.Body.Close()
			body, _ := io.ReadAll(res.Body)
			bodyStr := strings.Trim(string(body), "\n")
			logrus.Debugf("Body - Want: %s - Got: %v", *expectedBody, bodyStr)
			if bodyStr == *expectedBody {
				os.Exit(0)
			}
		} else {
			logrus.Infoln("Connection established!")
			os.Exit(0)
		}

		time.Sleep(sleepDuration)
		elapsed := time.Now().Sub(startTime)

		if elapsed > timeoutDuration {
			logrus.Errorf("Timeout")
			os.Exit(1)
		}
	}
}

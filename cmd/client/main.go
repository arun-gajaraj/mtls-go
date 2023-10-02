package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	l "github.com/sirupsen/logrus"
)

func main() {

	clientCert, err := tls.LoadX509KeyPair("client-cert.pem", "client-key.pem")
	if err != nil {
		l.WithError(err).Fatal("error loading client certificate")
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		InsecureSkipVerify: true, //
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		l.WithError(err).Errorln("error making client request")
		return
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		l.WithError(err).Errorln("error reading response body")
		return
	}

	fmt.Println(string(res))

}

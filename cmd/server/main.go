package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"

	l "github.com/sirupsen/logrus"
)

func main() {

	// Load Server certificate and key
	serverCert, err := tls.LoadX509KeyPair("server-cert.pem", "server-key.pem")
	if err != nil {
		l.WithError(err).Fatalln("error opening server cert and key")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    getClientCAs(),
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello mTLS World!")
	})

	// server cert already loaded in Server's tls config, hence blank string
	l.Fatal(server.ListenAndServeTLS("", ""))

}

func getClientCAs() *x509.CertPool {
	clientCACertPool := x509.NewCertPool()

	b, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		l.Errorln("error reading ca cert file", err)
	}

	block, _ := pem.Decode(b)
	if block == nil {
		l.Errorln("Failed to parse PEM certificate")
		return clientCACertPool
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		l.Errorln("error parsing ca certificate", err)
	}

	clientCACertPool.AddCert(cert)

	return clientCACertPool

}

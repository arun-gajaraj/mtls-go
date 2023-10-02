package main

import (
	"crypto/tls"
	"fmt"
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

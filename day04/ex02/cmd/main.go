package main

import (
	"OldCow/internal/router"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	CertPath            string = "../cert/candy.tld/cert.pem"
	KeyPath             string = "../cert/candy.tld/key.pem"
	RootCertificatePath string = "../cert/minica.pem"
)

func main() {
	// add endpoints
	router := router.SetUpRouter()

	// create a certificate pool and load all the CA certificates that you
	// want to validate a client against
	clientCA, err := os.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	clientCAPool := x509.NewCertPool()
	clientCAPool.AppendCertsFromPEM(clientCA)
	log.Println("ClientCA loaded")

	serverCerts, err := tls.LoadX509KeyPair(CertPath, KeyPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ServerCA loaded")

	// configure http server with tls configuration
	s := &http.Server{
		Handler: router,
		Addr:    ":3333",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{serverCerts},
			ClientCAs:    clientCAPool,
			ClientAuth:   tls.RequireAndVerifyClientCert,
		},
	}

	log.Println("starting server")

	// use server.ListenAndServeTLS instead of http.ListenAndServeTLS
	log.Fatal(s.ListenAndServeTLS("", ""))
}

func CertificateInfo(cert *x509.Certificate) string {
	if cert.Subject.CommonName == cert.Issuer.CommonName {
		return fmt.Sprintf("    Self-signed certificate %v\n", cert.Issuer.CommonName)
	}
	s := fmt.Sprintf("    Subject %v\n", cert.DNSNames)
	s += fmt.Sprintf("    Usage %v\n", cert.ExtKeyUsage)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.CommonName)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.SerialNumber)
	return s
}

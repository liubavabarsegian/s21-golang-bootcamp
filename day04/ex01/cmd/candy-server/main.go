package main

import (
	"BuyCandy/internal/router"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	CertPath            string = "../../cert/candy-server/cert.pem"
	KeyPath             string = "../../cert/candy-server/key.pem"
	RootCertificatePath string = "../../cert/minica.pem"
)

func main() {

	// add endpoints
	router := router.SetUpRouter()

	// create a certificate pool and load all the CA certificates that you
	// want to validate a client against
	clientCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	clientCAPool := x509.NewCertPool()
	clientCAPool.AppendCertsFromPEM(clientCA)
	log.Println("ClientCA loaded")

	// configure http server with tls configuration
	s := &http.Server{
		Handler: router,
		Addr:    ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:  clientCAPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			// Loads the server's certificate and sends it to the client
			GetCertificate: func(info *tls.ClientHelloInfo) (certificate *tls.Certificate, e error) {
				log.Println("client requested certificate")
				c, err := tls.LoadX509KeyPair(CertPath, KeyPath)
				if err != nil {
					fmt.Printf("Error loading server key pair: %v\n", err)
					return nil, err
				}
				return &c, nil
			},
			// Call back function to print client certificate details
			VerifyPeerCertificate: func(rawCerts [][]byte, chains [][]*x509.Certificate) error {
				if len(chains) > 0 {
					fmt.Println("Verified certificate chain from peer:")
					for _, v := range chains {
						for i, cert := range v {
							fmt.Printf("  Cert %d:\n", i)
							fmt.Printf(CertificateInfo(cert))
						}
					}
				}
				return nil
			},
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

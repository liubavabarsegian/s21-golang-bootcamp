package main

import (
	"BuyCandy/internal/api/request"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	RootCertificatePath string = "../../cert/minica.pem"
	ClientCertPath      string = "../../cert/candy-client/cert.pem"
	ClientKeyPath       string = "../../cert/candy-client/key.pem"
)

func main() {
	// pool for all the servers that you want to authenticate
	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")

	// configure TLS on http.Client
	c := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAPool,
				// Load clients key-pair. This will be sent to server
				GetClientCertificate: func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
					c, err := tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
					if err != nil {
						fmt.Printf("Error loading client key pair: %v\n", err)
						return nil, err
					}
					return &c, nil
				},
				// print  information about the certificate received from server
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
		},
	}

	// prepare a request
	u := url.URL{Scheme: "https", Host: "candy-server:8080", Path: "buy_candy"}
	requestBody := request.BuyCandyRequestBody{
		CandyCount: 2,
		CandyType:  "AA",
		Money:      31,
	}

	body, _ := json.Marshal(requestBody)
	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewReader(body))
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}

	// make the request
	var data string
	if data, err = callServer(c, r); err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}

func callServer(c http.Client, r *http.Request) (string, error) {
	response, err := c.Do(r)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// print the data
	return string(data), nil
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

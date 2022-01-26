package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {

	// Load CA 'Certificate Authority'
	caLoad, err := tls.LoadX509KeyPair("ca.crt", "ca.key")
	if err != nil {
		panic(err)
	}
	caCert, err := x509.ParseCertificate(caLoad.Certificate[0])
	if err != nil {
		panic(err)
	}

	// Prepare certificate
	certificate := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization:  []string{"www.awesomedogwebsitefordogs.com"},
			Country:       []string{"USA"},
			Province:      []string{"OHIO"},
			Locality:      []string{"COLUMBUS"},
			StreetAddress: []string{"NEIGHTBORHOOD DOGS"},
			PostalCode:    []string{"4322BARK"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1000, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, certificate, caCert, pub, caLoad.PrivateKey)
	if err != nil {
		log.Println("create certificate failed", err)
		return
	}

	// Public key
	certOut, err := os.Create("certificate.crt")
	if err != nil {
		log.Println("write public key failed", err)
		return
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	certOut.Close()
	log.Print("written certificate.crt\n")

	// Private key
	keyOut, err := os.OpenFile("certificate.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Println("write private key failed", err)
		return
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("written certificate.key\n")

}

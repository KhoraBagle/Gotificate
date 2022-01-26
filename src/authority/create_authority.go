package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"os"
	"time"
)

// I really wanted to do this via eliptic curves but that would be out of scope of the time I had to develop this.
// Additionally there is currently still limiteed browser and os support for ECC.
// However, considering that 256 bit ECC is 10,000 times stronger than 2048 RSA and is significatntly faster, it is the nest step
// in assymetric encryption and I would love to explore/develop it further.

func main() {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization:  []string{"DOG AUTHORITY"},
			Country:       []string{"USA"},
			Province:      []string{"OHIO"},
			Locality:      []string{"COLUMBUS"},
			StreetAddress: []string{"DOG PARK"},
			PostalCode:    []string{"43221"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1000, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
		return
	}

	// Public key
	certOut, err := os.Create("ca.crt")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	certOut.Close()
	log.Print("written cert.pem\n")

	// Private key
	keyOut, err := os.OpenFile("ca.key", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	log.Print("written key.pem\n")
}

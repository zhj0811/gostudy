package main

import (
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
	"time"

	keystore "github.com/pavel-v-chernykh/keystore-go/v3"
)

func main() {
	password := []byte("test1234")
	pke, _ := ioutil.ReadFile("./server.key")
	pce, _ := ioutil.ReadFile("./server.crt")
	p, _ := pem.Decode(pke)
	if p.Type != "PRIVATE KEY" {
		log.Fatal("Should be a rsa private key")
	}
	crt, _ := pem.Decode(pce)
	if crt.Type != "CERTIFICATE" {
		log.Fatal("Should be a Certificate")
	}

	ks := keystore.KeyStore{
		"server": &keystore.PrivateKeyEntry{
			Entry: keystore.Entry{
				CreationTime: time.Now(),
			},
			PrivateKey: p.Bytes,
			CertificateChain: []keystore.Certificate{
				{
					Type:    "X509",
					Content: crt.Bytes,
				},
			},
		},
	}
	f, err := os.Create("./server.keystore.jks")
	if err != nil {
		log.Fatal("create file failed")
	}
	err = keystore.Encode(f, ks, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success")

	cafile, _ := ioutil.ReadFile("./ca.crt")
	caBlock, _ := pem.Decode(cafile)
	truststore := keystore.KeyStore{
		"ca": &keystore.TrustedCertificateEntry{
			Entry: keystore.Entry{
				CreationTime: time.Now(),
			},
			Certificate: keystore.Certificate{
				Type:    "X509",
				Content: caBlock.Bytes,
			},
		},
	}
	tustFile, err := os.Create("./server.truststore.jks")
	if err != nil {
		log.Fatal("create tustFile failed")
	}
	err = keystore.Encode(tustFile, truststore, password)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Success Truststore")
}

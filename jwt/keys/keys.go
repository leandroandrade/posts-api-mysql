package keys

import (
	"crypto/rsa"
	"os"
	"bufio"
	"encoding/pem"
	"crypto/x509"
	"log"
)

var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func init() {
	loadPrivateKey()
	loadPublicKey()
}

func loadPrivateKey() {
	log.Println("loading RSA private key...")

	privateKeyFile, err := os.Open("/tmp/private_key")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	PrivateKey, err = x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		panic(err)
	}

	log.Println("RSA private key loaded successful!")
}

func loadPublicKey() {
	log.Println("loading RSA public key...")

	publicKeyFile, err := os.Open("/tmp/public_key.pub")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	pembytes := make([]byte, pemfileinfo.Size())

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	PublicKey = publicKeyImported.(*rsa.PublicKey)

	log.Println("RSA public key loaded successful!")
	//if !ok {
	//	panic(err)
	//}
	//PublicKey = rsaPub

}

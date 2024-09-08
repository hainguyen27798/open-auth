package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-open-auth/global"
	"os"
)

var (
	privateFileName = "private.pem"
	publicFileName  = "public.pem"
)

func GenerateRSA(bits int) {
	_, err := os.Stat(privateFileName)

	if os.IsNotExist(err) {
		privateKey, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			panic(err)
		}

		// Save the private key
		privateFile, err := os.Create(privateFileName)
		if err != nil {
			panic(err)
		}
		defer func(privateFile *os.File) {
			err := privateFile.Close()
			if err != nil {
				panic(err)
			}
		}(privateFile)

		privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privatePem := pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateBytes,
		}
		if err := pem.Encode(privateFile, &privatePem); err != nil {
			panic(err)
		}

		// Save the public key
		publicFile, err := os.Create(publicFileName)
		if err != nil {
			panic(err)
		}
		defer func(publicFile *os.File) {
			err := publicFile.Close()
			if err != nil {
				panic(err)
			}
		}(publicFile)

		publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			panic(err)
		}

		publicPem := pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicBytes,
		}
		if err := pem.Encode(publicFile, &publicPem); err != nil {
			panic(err)
		}

		fmt.Println("generate rsa")
	}

	privateKey, err := os.ReadFile(privateFileName)
	if err != nil {
		panic(err)
	}

	publicKey, err := os.ReadFile(publicFileName)
	if err != nil {
		panic(err)
	}

	if err := os.Setenv(global.TokenPrivateKey, string(privateKey)); err != nil {
		panic(err)
	}

	if err := os.Setenv(global.TokenPublicKey, string(publicKey)); err != nil {
		panic(err)
	}
}

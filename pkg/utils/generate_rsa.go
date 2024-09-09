package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/open-auth/global"
	"os"
	"strings"
)

var (
	privateSuffixFileName = "private.pem"
	publicSuffixFileName  = "public.pem"
	folder                = ".secrets"
)

func GenerateRSA(bits int, scope global.Scope) {
	privateFilePath := fmt.Sprintf("%s/%s_%s", folder, strings.ToLower(string(scope)), privateSuffixFileName)
	publicFilePath := fmt.Sprintf("%s/%s_%s", folder, strings.ToLower(string(scope)), publicSuffixFileName)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err := os.Mkdir(folder, os.ModePerm); err != nil {
			panic(err)
			return
		}
	}

	if _, err := os.Stat(privateFilePath); os.IsNotExist(err) {
		privateKey, err := rsa.GenerateKey(rand.Reader, bits)
		if err != nil {
			panic(err)
			return
		}

		// Save the private key
		privateFile, err := os.Create(privateFilePath)
		if err != nil {
			panic(err)
			return
		}
		defer func(privateFile *os.File) {
			err := privateFile.Close()
			if err != nil {
				panic(err)
				return
			}
		}(privateFile)

		privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privatePem := pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateBytes,
		}
		if err := pem.Encode(privateFile, &privatePem); err != nil {
			panic(err)
			return
		}

		// Save the public key
		publicFile, err := os.Create(publicFilePath)
		if err != nil {
			panic(err)
			return
		}
		defer func(publicFile *os.File) {
			err := publicFile.Close()
			if err != nil {
				panic(err)
				return
			}
		}(publicFile)

		publicBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
		if err != nil {
			panic(err)
			return
		}

		publicPem := pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: publicBytes,
		}
		if err := pem.Encode(publicFile, &publicPem); err != nil {
			panic(err)
			return
		}

		fmt.Println("generate rsa")
	}

	privateKey, err := os.ReadFile(privateFilePath)
	if err != nil {
		panic(err)
		return
	}

	publicKey, err := os.ReadFile(publicFilePath)
	if err != nil {
		panic(err)
		return
	}

	tokenPrivateKey := fmt.Sprintf("%s_%s", scope, global.TokenPrivateKey)
	if err := os.Setenv(tokenPrivateKey, string(privateKey)); err != nil {
		panic(err)
		return
	}

	tokenPublicKey := fmt.Sprintf("%s_%s", scope, global.TokenPublicKey)
	if err := os.Setenv(tokenPublicKey, string(publicKey)); err != nil {
		panic(err)
		return
	}
}

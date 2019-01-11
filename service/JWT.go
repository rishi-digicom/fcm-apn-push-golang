package service

import (
	"bufio"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GenerateJwtToken(key_id string, issuer_claim_key string, file_path string) (string, error) {

	alg := "ES256"
	kid := key_id

	iat := time.Now().Unix()
	iss := issuer_claim_key

	token := jwt.New(jwt.SigningMethodES256)
	token.Claims = jwt.MapClaims{
		"iss": iss,
		"iat": iat,
	}

	token.Header["alg"] = alg
	token.Header["kid"] = kid

	privateKey, err := importP8File(file_path)

	if err != nil {
		return "", err
	}

	return token.SignedString(privateKey)
}

func importP8File(file_path string) (interface{}, error) {

	privateKeyFile, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p8FileInfo, _ := privateKeyFile.Stat()

	var size int64 = p8FileInfo.Size()

	// assigning memory of its size to store bytes
	p8Bytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)

	_, err = buffer.Read(p8Bytes)
	if err != nil {
		return nil, err
	}

	data, _ := pem.Decode([]byte(p8Bytes))

	defer privateKeyFile.Close()

	privateKey, err := x509.ParsePKCS8PrivateKey(data.Bytes)

	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

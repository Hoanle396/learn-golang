package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var (
	key   []byte = []byte("5e2aef5769b4241d5e2aef5769b4241d")
	nonce []byte = []byte("9123ace458f3")
)

func Encrypt(citizenID string) (string, error) {
	// create a AES key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// use the Galois Counter Mode (GCM) 128-bit, block cipher with a standard nonce
	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext := []byte(citizenID)

	// encrypt
	cipherText := aesGcm.Seal(nil, nonce, plaintext, nil)
	encryptedStr := base64.StdEncoding.EncodeToString(cipherText)
	return encryptedStr, nil
}
func Decrypt(encryptedStr string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesGcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

func main() {
	encrypt, err := Encrypt("1309913659936")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(encrypt)

	plaintext, err := Decrypt(encrypt)
	if err != nil {
		fmt.Printf("Error %s ", err.Error())
	}
	fmt.Printf("[AES_GCM decrypt] ==> %s ", string(plaintext))
}

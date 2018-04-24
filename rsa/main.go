package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func generateRSAKeyPair(bit int) (*rsa.PrivateKey, *rsa.PublicKey) {
	priv, err := rsa.GenerateKey(rand.Reader, bit)
	if err != nil {
		fmt.Println(err.Error())
	}

	pub := &priv.PublicKey

	return priv, pub
}

func encryptMessage(msg []byte, label []byte, pub *rsa.PublicKey) []byte {
	hash := sha256.New()
	cipherText, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cipherText
}

func decryptMessage(cipherText []byte, label []byte, priv *rsa.PrivateKey) []byte {
	hash := sha256.New()
	label = []byte("")
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, priv, cipherText, label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return plainText
}

func main() {
	// Generate RSA Keys
	alicePrivateKey, alicePublicKey := generateRSAKeyPair(2048)

	// Encrypt Message
	msg := []byte("The code must be like a piece of music")
	label := []byte("")
	cipherText := encryptMessage(msg, label, alicePublicKey)

	fmt.Printf("OAEP encrypted [%s] to \n[%x]\n", string(msg), cipherText)
	fmt.Println()

	// Decrypt Message
	plainText := decryptMessage(cipherText, label, alicePrivateKey)
	fmt.Printf("OAEP decrypted [%x] to \n[%s]\n", cipherText, string(plainText))
}

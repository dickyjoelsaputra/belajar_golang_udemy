package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	plainText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	key := []byte("my 32 length key................")
	iv := make([]byte, aes.BlockSize)

	// Enkripsi pesan

	encryptedBase64, err := encryptAES(plainText, key, iv)
	if err != nil {
		log.Println("Gagal mengenkripsi pesan:", err)
		return
	}

	fmt.Println("Hasil Enkripsi:", encryptedBase64)

	// Dekripsi pesan

	// encryptedBase64 := "QK0IcGNYK7Pu/p0XTbBrPPOdrllEI1+4IrZnyHX2k4s="

	decrypted, err := decryptAES(encryptedBase64, key, iv)
	if err != nil {
		fmt.Println("Gagal mendekripsi pesan terenkripsi:", err)
		return
	}

	fmt.Println("Hasil Dekripsi:", decrypted)
}

func decryptAES(encryptedBase64 string, key, iv []byte) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encrypted))
	mode.CryptBlocks(decrypted, encrypted)

	decrypted = unpad(decrypted)
	return string(decrypted), nil
}

func unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func encryptAES(plainText string, key, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	paddedText := pad([]byte(plainText))

	mode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(paddedText))
	mode.CryptBlocks(encrypted, paddedText)

	// Encode hasil enkripsi dalam base64
	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
	return encryptedBase64, nil
}

func pad(data []byte) []byte {
	blockSize := aes.BlockSize
	padding := blockSize - (len(data) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

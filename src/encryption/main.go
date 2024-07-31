package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/nietthijmen/tavern/src/config"
	"github.com/rs/zerolog/log"
	"io"
)

func getKey() []byte {
	key := config.ReadEnv("ENCRYPTION_KEY", "default")
	decodedData, err := base64.StdEncoding.DecodeString(key)

	if err != nil {
		log.Panic().Err(err).Msg("Failed to decode encryption key")
	}

	return decodedData
}

func Encrypt(data string) (string, error) {
	block, err := aes.NewCipher(getKey())
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encodedValue string) (string, error) {
	block, err := aes.NewCipher(getKey())
	if err != nil {
		return "", err
	}

	ciphertextBytes, err := base64.URLEncoding.DecodeString(encodedValue)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertext := ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

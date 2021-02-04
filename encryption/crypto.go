package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/Luzifer/go-openssl"
)

// CreateHash creates a hash of the given key
func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// Decrypt decrypts the given data with the given passphrase
func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Printf("Decrypt 40: %v", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Printf("Decrypt 44: %v", err)
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Printf("Decrypt 51: %v", err)
	}
	return plaintext
}

// DecryptMessage to test a solution This works when recieving the encrypted string from JSON
func DecryptMessage(key, data string) string {
	encrypted := data
	secret := key

	o := openssl.New()

	dec, err := o.DecryptBytes(secret, []byte(encrypted))
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}

	return string(dec)
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext := encrypt([]byte("Enter Message Here!!!!!"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext := Decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
}

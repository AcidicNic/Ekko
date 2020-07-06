package main

import "github.com/AcidicNic/Ekko/src/server"

func main() {
	server.Start()
	// data := []byte("U2FsdGVkX1+Ck/su6DsLqgMNuXuydGqjtYs+ceLY13Y=")

	// block, err := aes.NewCipher([]byte(encryption.CreateHash("omar")))
	// if err != nil {
	// 	panic(err)
	// }
	// gcm, err := cipher.NewGCM(block)
	// if err != nil {
	// 	log.Printf("Decrypt 44: %v", err)
	// }
	// nonceSize := gcm.NonceSize()
	// nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	// plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	// if err != nil {
	// 	log.Printf("Decrypt 51: %v", err)
	// }
	// fmt.Println(plaintext)
}

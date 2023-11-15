// This program extracts a secret from HashiCorp vault and performs Encryption using that key
// Some of the values are hardcoded as i was getting issues with VSCode

package main

import (
	//"bufio"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	//"os"

	vault "github.com/hashicorp/vault/api"
)

// This is the accompanying code for the Developer Quick Start.
// WARNING: Using root tokens is insecure and should never be done in production!
func main() {
	config := vault.DefaultConfig()

	config.Address = "http://127.0.0.1:8200"

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("unable to initialize Vault client: %v", err)
	}

	//Reading from cmd line was not working with VSCode
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Vault Token: ")
	token1, _ := reader.ReadString('\n')
	fmt.Println("Token " + token1)*/

	//Do not hard code values, i was getting VS Code error, so hardcoding the value
	//Authenticate
	client.SetToken("hvs.Fpap2r1qrQ9GE9z6SyUwH5kp")

	fmt.Print(" token set ")

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("GoLangExample").Get(context.Background(), "GoLangEncryptionKey")
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}
	fmt.Print(secret)
	secretvalue, ok := secret.Data["GoLangEncryptionKey1"].(string)
	if !ok {
		log.Fatalf("value type assertion failed: %T %#v", secret.Data["password"], secret.Data["password"])
	}

	fmt.Print("Secret Extracted from Vault " + secretvalue)

	cipherkey, err := aes.NewCipher([]byte(secretvalue))
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	stringToBeEncrypted := []byte("String to be encrypted")

	cipherText := make([]byte, aes.BlockSize+len(stringToBeEncrypted))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	fmt.Print(string(iv))

	stream := cipher.NewCFBEncrypter(cipherkey, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], stringToBeEncrypted)

	fmt.Print("Encrypted message " + string(cipherText))
}

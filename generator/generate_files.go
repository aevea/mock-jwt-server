package generator

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

const PrivateKeyName = "mock-jwt-private.pem"
const PublicKeyName = "mock-jwt-public.pem"

func (g *Generator) GenerateFiles(path string) error {

	// Generate a new RSA private key with 2048 bits
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	// Encode the private key to the PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	privateKeyPath := filepath.Join(path, PrivateKeyName)
	privateKeyFile, err := os.Create(privateKeyPath)
	if err != nil {
		return err
	}
	pem.Encode(privateKeyFile, privateKeyPEM)
	privateKeyFile.Close()

	// Extract the public key from the private key
	publicKey := &privateKey.PublicKey

	// Encode the public key to the PEM format
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	publicKeyPath := filepath.Join(path, PublicKeyName)
	publicKeyFile, err := os.Create(publicKeyPath)
	if err != nil {
		return err
	}
	pem.Encode(publicKeyFile, publicKeyPEM)
	publicKeyFile.Close()

	fmt.Println("RSA key pair generated successfully!")

	g.privateKey = privateKey
	g.publicKey = publicKey

	return nil
}

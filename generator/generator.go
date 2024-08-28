package generator

import (
	"crypto/rsa"
)

type Generator struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func InitGenerator(path string) (*Generator, error) {
	gen := &Generator{}

	err := gen.GenerateFiles(path)

	if err != nil {
		return nil, err
	}

	return gen, nil
}

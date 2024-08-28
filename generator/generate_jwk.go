package generator

import (
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

type KeyReturn struct {
	Keys []jwk.Key `json:"keys"`
}

func (g *Generator) GenerateJWK() (string, error) {

	key, err := jwk.FromRaw(g.publicKey)

	if err != nil {
		return "", err
	}

	key.Set("alg", "RS256")
	key.Set("kid", "mykey")

	keys := KeyReturn{
		Keys: []jwk.Key{
			key,
		},
	}

	data, err := json.Marshal(keys)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

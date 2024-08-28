package generator

import "github.com/golang-jwt/jwt/v5"

func (g *Generator) GenerateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":      "tester",
		"org_code": "test_org",
		"org_name": "Test Org",
	})

	token.Header["kid"] = "mykey"

	s, err := token.SignedString(g.privateKey)

	if err != nil {
		return "", err
	}

	return s, nil
}

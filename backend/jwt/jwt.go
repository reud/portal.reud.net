package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"log"
	"os"
)

func Validate(key string) (jwtMap *map[string]*json.RawMessage, err error) {
	set, err := jwk.Fetch("https://reud.auth0.com/.well-known/jwks.json")
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return nil, err
	}

	p, err := jws.VerifyWithJWKSet([]byte(key), set, nil)
	fmt.Println(string(p))

	var jwt map[string]*json.RawMessage
	err = json.Unmarshal(p, &jwt)
	if err != nil {
		return nil, err
	}

	sub := os.Getenv("SUB_TOKEN")

	if sub != string(*jwt["sub"]) {
		return nil, errors.New("u r not reud")
	}

	fmt.Println(string(*jwt["sub"]))

	return &jwt, err
}

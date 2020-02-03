package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"log"
	"os"
	"strings"
)

func Env_load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
}

func Validate(key string) (jwtMap *map[string]*json.RawMessage, err error) {
	Env_load()
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

	exsub := os.Getenv("SUB_TOKEN")
	sub := strings.Replace(string(*jwt["sub"]), "\"", "", -1)

	if exsub != sub {
		return nil, errors.New("u r not reud")
	}

	fmt.Println(string(*jwt["sub"]))

	return &jwt, err
}

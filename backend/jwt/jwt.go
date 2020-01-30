package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"log"
)

func Validate(key string) (jwtMap *map[string]*json.RawMessage, err error) {
	set, err := jwk.Fetch("https://reud.auth0.com/.well-known/jwks.json")
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return nil, err
	}

	// TODO: 多分ここ今全アプリケーション<->全API　に対して許可になっているのでいい感じにする。 俺のsubクレームと比較でいい気がする
	p, err := jws.VerifyWithJWKSet([]byte(key), set, nil)
	fmt.Println(string(p))

	var jwt map[string]*json.RawMessage
	err = json.Unmarshal(p, &jwt)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(*jwt["sub"]))

	return &jwt, err
}

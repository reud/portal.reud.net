package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"log"
)

type T struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Iat int    `json:"iat"`
	Exp int    `json:"exp"`
	Azp string `json:"azp"`
	Gty string `json:"gty"`
}

func Validate(key string) (payload *T, err error) {
	set, err := jwk.Fetch("https://reud.auth0.com/.well-known/jwks.json")
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return nil, err
	}

	// TODO: 多分ここ今全アプリケーション<->全API　に対して許可になっているのでいい感じにする。 俺のsubクレームと比較でいい気がする
	p, err := jws.VerifyWithJWKSet([]byte(key), set, nil)
	data := &T{}
	err = json.Unmarshal(p, data)
	if err != nil {
		return nil, err
	}

	return data, err
}

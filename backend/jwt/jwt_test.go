package jwt

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

// 雑なmiddleJWT確認用テスト
func TestValidation(t *testing.T) {
	// TODO: .envで設定する。
	key := "いいヵんじのJWT"
	result, err := Validate(key)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Fatalf("result: \n%#v", result.Sub)

}

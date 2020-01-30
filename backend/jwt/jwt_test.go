package jwt

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func Env_load() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file %s", err.Error())
	}
}

// 雑なmiddleJWT確認用テスト
func TestValidation(t *testing.T) {
	Env_load()
	// TODO: .envで設定する。
	key := os.Getenv("ACCESS_TOKEN")
	result, err := Validate(key)
	if err != nil {
		t.Fatalf("%#v", err)
	}

	t.Fatalf("result: \n%#v", string(*(*result)["sub"]))

}

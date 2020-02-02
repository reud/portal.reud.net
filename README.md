reud.portal.net
===

# TODOこれはもう古いので書き直す(swagger-codegen-cli -> openapi-generator)
reudのポートフォリオ

# contains
- Auth0
- Vue.js + Typescript
- Go


# 環境構築(Mac)
自分用の開発環境構築です。

## Codegen
- サーバとフロントの通信を円滑にするために、swagger.yamlから、サーバ用コード、クライアント側用コードの生成をしています。

- サーバは[go-swagger](https://github.com/go-swagger/go-swagger),フロントは[swagger-codegen-cli v2](https://github.com/swagger-api/swagger-codegen)を使用しています。

### go-swaggerのインストール

```bash
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```

### [Docker]swagger-codegen-cli(v2系)のインストール
Dockerが一番楽なのでDockerを使用することをお勧めします。

```bash
$ docker pull swaggerapi/swagger-codegen-cli
```


# 参考
- [go-swagger を使って swagger から Goのサーバとクライアントのコードを生成する - Qiita](https://qiita.com/o_tyazuke/items/43bd362e8e427aa0e340)
- [swagger-api/swagger-codegen: swagger-codegen contains a template-driven engine to generate documentation, API clients and server stubs in different languages by parsing your OpenAPI / Swagger definition.](https://github.com/swagger-api/swagger-codegen#swagger-codegen-cli-docker-image)


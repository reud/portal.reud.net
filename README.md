reud.portal.net
===

reudのポートフォリオ

# contains
- Auth0
- Vue.js + Typescript
- Go


# 環境構築(Mac)
自分用の開発環境構築です。

## Codegen
- サーバとフロントの通信を円滑にするために、swagger.yamlから、サーバ用コード、クライアント側用コードの生成をしています。

- サーバは[go-swagger](https://github.com/go-swagger/go-swagger),フロントは[openapi-generator](https://github.com/OpenAPITools/openapi-generator)を使用しています。

### go-swaggerのインストール

```bash
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```

### openapi-generatorのインストール

```bash
$ brew install openapi-generator
```

# TODO
- [ ] CORS周り見直す
- [ ] フロントからdelete出来る様にする。
- [ ] Azure VMからAlibaba cloudに乗り換える
- [ ] databaseをどうにかする
- [ ] 画面をもう少し考える
- [ ] 本棚のタイトルを入力できる様にする
- [ ] 諸々リファクタリング
- [ ] 名前変更機能の実装

# 参考
- [go-swagger を使って swagger から Goのサーバとクライアントのコードを生成する - Qiita](https://qiita.com/o_tyazuke/items/43bd362e8e427aa0e340)
- [OpenAPITools/openapi-generator: OpenAPI Generator allows generation of API client libraries (SDK generation), server stubs, documentation and configuration automatically given an OpenAPI Spec (v2, v3)](https://github.com/OpenAPITools/openapi-generator#15---homebrew)

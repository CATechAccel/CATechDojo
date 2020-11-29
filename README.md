# CATechDojo
CA Tech Dojo: https://techbowl.co.jp/techtrain/missions/12

## API仕様書
1. https://editor.swagger.io/ を開く
2. https://github.com/CATechAccel/CATechDojo/blob/main/swagger.yml の内容をコピーして貼り付ける

▼swagger.ymlの内容
``` yml
swagger: "2.0"
info:
  description: "TechTrain MISSION ゲームAPI入門仕様\n
  まずはこのAPI仕様に沿って機能を実装してみましょう。\n
  \n
  この画面の各APIの「Try it out」->「Execute」を利用することで\n
  ローカル環境で起動中のAPIにAPIリクエストをすることができます。"
  version: "1.0.0"
  title: "TechTrain MISSION Game API"
host: "localhost:8080"
tags:
  - name: "user"
    description: "ユーザ関連API"
  - name: "gacha"
    description: "ガチャ関連API"
  - name: "character"
    description: "キャラクター関連API"
schemes:
  - "http"
paths:
  /user/create:
    post:
      tags:
        - "user"
      summary: "ユーザ情報作成API"
      description: "ユーザ情報を作成します。\n
      ユーザの名前情報をリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。"
      consumes:
        - "application/json"
      produces:
        - "application/json"
      parameters:
        - in: "body"
          name: "body"
          description: "Request Body"
          required: true
          schema:
            $ref: "#/definitions/UserCreateRequest"
      responses:
        200:
          "description": "A successful response."
          "schema":
            "$ref": "#/definitions/UserCreateResponse"

  /user/get:
    get:
      tags:
        - "user"
      summary: "ユーザ情報取得API"
      description: "ユーザ情報を取得します。\n
      ユーザの認証と特定の処理はリクエストヘッダのx-tokenを読み取ってデータベースに照会をします。"
      consumes:
        - "application/json"
      produces:
        - "application/json"
      parameters:
        - in: "header"
          name: "x-token"
          description: "認証トークン"
          required: true
          type: "string"
      responses:
        200:
          "description": "A successful response."
          "schema":
            "$ref": "#/definitions/UserGetResponse"

  /user/update:
    put:
      tags:
        - "user"
      summary: "ユーザ情報更新API"
      description: "ユーザ情報の更新をします。\n
      初期実装では名前の更新を行います。"
      consumes:
        - "application/json"
      produces:
        - "application/json"
      parameters:
        - in: "header"
          name: "x-token"
          description: "認証トークン"
          required: true
          type: "string"
        - in: "body"
          name: "body"
          description: "Request Body"
          required: true
          schema:
            $ref: "#/definitions/UserUpdateRequest"
      responses:
        200:
          "description": "A successful response."

  /gacha/draw:
    post:
      tags:
        - "gacha"
      summary: "ガチャ実行API"
      description: "ガチャを引いてキャラクターを取得する処理を実装します。\n
      獲得したキャラクターはユーザ所持キャラクターテーブルへ保存します。\n
      同じ種類のキャラクターでもユーザは複数所持することができます。\n
      \n
      キャラクターの確率は等倍ではなく、任意に変更できるようテーブルを設計しましょう。"
      consumes:
        - "application/json"
      produces:
        - "application/json"
      parameters:
        - in: "header"
          name: "x-token"
          description: "認証トークン"
          required: true
          type: "string"
        - in: "body"
          name: "body"
          description: "Request Body"
          required: true
          schema:
            $ref: "#/definitions/GachaDrawRequest"
      responses:
        200:
          "description": "A successful response."
          "schema":
            "$ref": "#/definitions/GachaDrawResponse"

  /character/list:
    get:
      tags:
        - "character"
      summary: "ユーザ所持キャラクター一覧取得API"
      description: "ユーザが所持しているキャラクター一覧情報を取得します。"
      consumes:
        - "application/json"
      produces:
        - "application/json"
      parameters:
        - in: "header"
          name: "x-token"
          description: "認証トークン"
          required: true
          type: "string"
      responses:
        200:
          "description": "A successful response."
          "schema":
            "$ref": "#/definitions/CharacterListResponse"

definitions:
  UserCreateRequest:
    type: "object"
    properties:
      name:
        type: "string"
        description: "ユーザ名"
  UserCreateResponse:
    type: "object"
    properties:
      token:
        type: "string"
        description: "クライアント側で保存するトークン"
  UserGetResponse:
    type: "object"
    properties:
      name:
        type: "string"
        description: "ユーザ名"
  UserUpdateRequest:
    type: "object"
    properties:
      name:
        type: "string"
        description: "ユーザ名"
  GachaDrawRequest:
    type: "object"
    properties:
      times:
        type: "integer"
        description: "実行回数"
  GachaDrawResponse:
    type: "object"
    properties:
      results:
        type: "array"
        items:
          $ref: "#/definitions/GachaResult"
  GachaResult:
    type: "object"
    properties:
      characterID:
        type: "string"
        description: "キャラクターID"
      name:
        type: "string"
        description: "キャラクター名"
  CharacterListResponse:
    type: "object"
    properties:
      characters:
        type: "array"
        items:
          $ref: "#/definitions/UserCharacter"
  UserCharacter:
    type: "object"
    properties:
      userCharacterID:
        type: "string"
        description: "ユニークID"
      characterID:
        type: "string"
        description: "キャラクターID"
      name:
        type: "string"
        description: "キャラクター名"
```

## github の運用ルール

このプロジェクトは、`IDD(Issue Driven Develop)`で行います！

### IDD(Issue Driven Develop)とは？

issue に対する開発を行い、必ず PullRequest も issue に向けた変更になっているように設計する開発手法です。

### ルール

1. main ブランチ への直接 Push は 🆖
2. ブランチルールに則った Branch 名で作業ブランチを切ること
3. PullRequest は必ず main ブランチ に向けて作ること

### ブランチルール

```
feat/[issue番号]/[issueの内容（英語で）]

[例]
feat/1/createBaseWebApplication
```

### 作業の流れ

1. [Project](https://github.com/CATechAccel/CATechDojo/projects/1)を確認して自分がアサインされている issue がないか確認
2. アサインされている issue の中から実装しようと決めた issue を「In Progress」に移動する
3. main ブランチをチェックアウトして、`git pull`する
4. ブランチルールに従い、作業用のブランチを新規作成
5. issue 内容を満たすようにローカル環境で開発を進める
   1. commit はできるだけ作業内容が後から追ったときにわかりやすい単位で細かく切ること！
   2. コミットメッセージは日本語で良いので丁寧に書くこと！
6. 開発が完了したら、`main`←`作業ブランチ`の PullRequest を作成する
7. Pull Request に`close #1`のように close コマンドを入力して、Pull Request をマージしたら紐づく issue が close されるようにしておく
8. Reviewers, Assignees, Labels の項目を issue と合わせて設定しておく
9. 「Files changed」で自分の書いたコードにバグなどがないか一通りチェックする
10. Reviewer に設定した人に Slack の`#notify-github`でレビュー依頼をする（レビュー依頼時は Pull Request のリンクも添付しましょう）

# Zenn投稿内容

# 概要
main.goに全ての責務が集約されたTodoアプリをクリーンアーキテクチャに修正していきます。
Part1で作成したTodoアプリを修正するので、簡単に[前回の記事](https://zenn.dev/kyuki/articles/806142fed1cc06)に目を通しておくと理解しやすいかもしれません。

## 完成系の構成
```shell
.
├── db
│   └── my.cnf
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
└── src
    ├── usecase # Application BusinessRules
    │   └── services # ビジネスロジックの実行手段を集約
    │       └── todo.go
    ├── domain # Enterprise Business Rules
    │   ├── models # DBのデーブルモデル
    │   │   └── todo.go
    │   └── repositories # repositoryのインターフェース
    │       └── todo.go
    ├── infra # Frameworks & Drivers層
    │   ├── database # DBの設定ファイル管理、接続、repositoryファイル
    │   │   ├── connection.go
    │   │   └── repositories
    │   │       └── todo.go
    │   └── http # クライアント側。htmlや遷移
    │       ├── public
    │       │   ├── edit.html
    │       │   └── index.html
    │       └── routes
    │           ├── page.go
    │           └── todo.go
    └── interface # interface Adapter層
        └── controllers # リクエスト内容の解析
            └── todo.go
```
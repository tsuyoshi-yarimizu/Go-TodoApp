# TODO APP
## 概要
簡易的なTODOアプリケーションです.

## 開発環境・言語等
### OS
macOS Mojave (ver 10.14.6)

### 開発言語
Go (ver 1.13)

### DB
Mysql (ver 5.7)

## 機能
|#|機能|URI|
|:---:|:---|:---|
|1|タスク一覧表示|GET /|
|2|タスク新規登録|POST /api/todo|
|3|タスク完了|POST /api/todo/:id/complete|
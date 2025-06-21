# Gin Practice

Go言語のGinフレームワークを使用したRESTful APIサーバーの実装例です。クリーンアーキテクチャを採用し、商品在庫管理システムのバックエンドAPIを提供します。

## 機能

- **マルチテナント対応**: 複数の店舗/拠点での在庫管理
- **商品管理**: 商品の登録・更新・削除
- **在庫管理**: 在庫数の管理と購入処理
- **ユーザー管理**: ユーザーの登録・更新・削除
- **ソフトデリート**: 全てのエンティティで論理削除をサポート

## 技術スタック

- **言語**: Go 1.22
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL 13
- **コンテナ**: Docker & Docker Compose
- **ロガー**: Zap

## アーキテクチャ

クリーンアーキテクチャに基づいた4層構造を採用しています：

```
┌─────────────────┐
│   Controller    │  HTTPリクエスト/レスポンス処理
├─────────────────┤
│    UseCase      │  ビジネスロジック
├─────────────────┤
│    Service      │  データアクセス層
├─────────────────┤
│     Model       │  ドメインモデル
└─────────────────┘
```

## プロジェクト構造

```
gin-practice/
├── app/
│   ├── config/          # 設定管理
│   ├── controller/      # HTTPハンドラー
│   ├── errors/          # カスタムエラー定義
│   ├── infra/           # インフラストラクチャ層
│   │   └── seed/        # シードデータ
│   ├── middleware/      # ミドルウェア
│   ├── migrate/         # マイグレーション
│   ├── model/           # ドメインモデル
│   ├── routes/          # ルーティング定義
│   ├── service/         # データアクセス層
│   └── usecase/         # ビジネスロジック層
├── main.go              # エントリーポイント
├── Dockerfile           # Dockerイメージ定義
├── docker-compose.yml   # Docker Compose設定
├── go.mod               # Goモジュール定義
└── go.sum               # 依存関係チェックサム
```

## セットアップ

### 前提条件

- Docker
- Docker Compose
- Go 1.22以降（ローカル開発の場合）

### インストール

1. リポジトリのクローン

```bash
git clone <repository-url>
cd gin-practice
```

2. Docker Composeで起動

```bash
# ビルドして起動
docker-compose up --build

# バックグラウンドで起動
docker-compose up -d
```

3. アプリケーションの確認

```bash
# ヘルスチェック
curl http://localhost:8080/ping
```

## 開発

### ローカルでの実行

```bash
# 依存関係のインストール
go mod download

# アプリケーションの実行
go run main.go
```

### 環境変数

| 変数名 | デフォルト値 | 説明 |
|--------|------------|------|
| DB_HOST | db | データベースホスト |
| DB_PORT | 5432 | データベースポート |
| DB_USER | user | データベースユーザー |
| DB_PASSWORD | password | データベースパスワード |
| DB_NAME | myapp | データベース名 |
| SERVER_PORT | 8080 | サーバーポート |
| ENV | development | 実行環境 |

### データベースマイグレーション

```bash
# Docker経由でマイグレーション実行
docker-compose run migrate

# 直接実行
go run app/migrate/migrate.go
```

## API エンドポイント

### ユーザー管理

- `GET /users` - ユーザー一覧取得
- `GET /users/:id` - ユーザー詳細取得
- `POST /users` - ユーザー作成
- `PUT /users/:id` - ユーザー更新
- `DELETE /users/:id` - ユーザー削除

### 商品管理

- `GET /products` - 商品一覧取得
- `POST /products` - 商品作成
- `PUT /products/:id` - 商品更新
- `DELETE /products/:id` - 商品削除

### 在庫管理

- `GET /inventories` - 在庫一覧取得
- `POST /inventories` - 在庫作成
- `POST /inventories/purchase` - 商品購入（在庫減少）
- `DELETE /inventories/:id` - 在庫削除

### テナント管理

- `GET /tenants` - テナント一覧取得
- `POST /tenants` - テナント作成
- `PUT /tenants/:id` - テナント更新
- `DELETE /tenants/:id` - テナント削除

### ヘルスチェック

- `GET /ping` - サーバー状態確認

## API使用例

### ユーザー作成

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山田太郎",
    "email": "yamada@example.com"
  }'
```

### 商品作成

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "ノートPC",
    "price": 98000
  }'
```

### 在庫登録

```bash
curl -X POST http://localhost:8080/inventories \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "tenant_id": 1,
    "quantity": 50
  }'
```

### 商品購入

```bash
curl -X POST http://localhost:8080/inventories/purchase \
  -H "Content-Type: application/json" \
  -d '{
    "product_id": 1,
    "tenant_id": 1,
    "quantity": 3
  }'
```

## Docker コマンド

```bash
# コンテナの状態確認
docker-compose ps

# ログの確認
docker-compose logs -f app

# データベースログの確認
docker-compose logs -f db

# コンテナの停止
docker-compose down

# ボリュームも含めて削除
docker-compose down -v

# コンテナへのアクセス
docker-compose exec app /bin/sh
```

## トラブルシューティング

### ポート競合エラー

ポート8080または5432が既に使用されている場合は、`docker-compose.yml`のポート設定を変更してください。

### データベース接続エラー

環境変数が正しく設定されているか確認してください。Docker Compose使用時は`DB_HOST=db`である必要があります。

### マイグレーションエラー

データベースが起動していることを確認してから、マイグレーションを実行してください。

# Gin Practice

Go言語のGinフレームワークを使用したRESTful APIサーバーの実装例です。クリーンアーキテクチャを採用し、商品在庫管理システムのバックエンドAPIを提供します。

## 機能

- **マルチテナント対応**: 複数の店舗/拠点での在庫管理
- **商品管理**: 商品の登録・更新・削除（バーコード対応）
- **在庫管理**: 在庫数の管理と入荷処理
- **価格設定**: 店舗別価格とセール機能（期間限定価格）
- **注文管理**: 商品購入と注文履歴管理
- **ユーザー管理**: ユーザーの登録・更新・削除
- **APIドキュメント**: Swagger UI による対話的なAPI仕様書

## 技術スタック

- **言語**: Go 1.22
- **Webフレームワーク**: Gin
- **ORM**: GORM
- **データベース**: PostgreSQL 13
- **コンテナ**: Docker & Docker Compose
- **APIドキュメント**: Swagger/OpenAPI
- **開発ツール**: Makefile, golangci-lint

## クイックスタート

### 前提条件

- Docker Desktop がインストールされていること
- Make コマンドが使用できること（macOS/Linuxは標準搭載）

### セットアップ

```bash
# リポジトリのクローン
git clone https://github.com/yourusername/gin-practice.git
cd gin-practice

# 初回起動（DBリセット + アプリケーション起動）
make reset

# 開発モード（ログ表示付き）で起動
make dev
```

アプリケーションは以下のURLでアクセスできます：
- API: http://localhost:8080
- Swagger UI: http://localhost:8080/swagger/index.html

## 主要なMakeコマンド

```bash
make help         # 利用可能なコマンド一覧を表示
make dev          # 開発モードで起動（ログ表示）
make test-v       # テストを詳細モードで実行
make swagger      # Swaggerドキュメントを生成
make db-shell     # PostgreSQLシェルに接続
make clean        # コンテナとデータを削除
```

詳細は `make help` を実行してください。

## アーキテクチャ

クリーンアーキテクチャに基づいた6層構造を採用しています：

```
┌─────────────────┐
│     Router      │  URLルーティング、ミドルウェア
├─────────────────┤
│   Controller    │  HTTPリクエスト/レスポンス処理
├─────────────────┤
│    UseCase      │  ビジネスロジック
├─────────────────┤
│   Validation    │  入力検証ロジック
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
│   ├── service/         # サービス層（データアクセス）
│   └── usecase/         # ユースケース層（ビジネスロジック）
│       ├── request/     # リクエストDTO
│       └── validation/  # バリデーションロジック
├── docs/                # Swaggerドキュメント
├── docker-compose.yml   # Docker Compose設定
├── Dockerfile          # Dockerイメージ定義
├── Makefile            # 開発用コマンド定義
├── go.mod              # Goモジュール定義
└── main.go             # エントリーポイント
```

## API エンドポイント

### ユーザー管理
- `GET /users` - ユーザー一覧取得
- `GET /users/:id` - ユーザー詳細取得
- `POST /users` - ユーザー作成
- `PUT /users/:id` - ユーザー更新
- `DELETE /users/:id` - ユーザー削除
- `GET /users/:id/orders` - ユーザーの注文履歴取得

### 商品管理
- `GET /products` - 商品一覧取得
- `POST /products` - 商品作成
- `PUT /products/:id` - 商品更新
- `DELETE /products/:id` - 商品削除

### 在庫管理
- `GET /inventories` - 在庫一覧取得
- `POST /inventories` - 在庫作成
- `PUT /inventories/:id` - 在庫更新
- `DELETE /inventories/:id` - 在庫削除
- `POST /inventories/restock` - 在庫入荷

### 価格設定
- `POST /inventories/:id/prices` - 価格設定作成
- `GET /inventories/:id/prices` - 価格履歴取得
- `GET /inventories/:id/prices/current` - 現在価格取得
- `PUT /inventories/:id/prices/:price_id` - 価格設定更新
- `DELETE /inventories/:id/prices/:price_id` - 価格設定削除

### 店舗管理
- `GET /tenants` - 店舗一覧取得
- `GET /tenants/:id` - 店舗詳細取得
- `POST /tenants` - 店舗作成
- `PUT /tenants/:id` - 店舗更新
- `DELETE /tenants/:id` - 店舗削除

### 注文管理
- `GET /orders` - 注文一覧取得
- `GET /orders/:id` - 注文詳細取得
- `POST /orders` - 注文作成
- `POST /orders/:id/cancel` - 注文キャンセル

詳細なAPI仕様は Swagger UI (http://localhost:8080/swagger/index.html) で確認できます。

## 開発

### 新機能の追加手順

1. モデルを作成 (`app/model/`)
2. `base_model.go` の `GetModels()` に追加
3. サービスを作成 (`app/service/`)
4. ユースケースを作成 (`app/usecase/`)
5. バリデーションを作成 (`app/usecase/validation/`)
6. コントローラーを作成 (`app/controller/`)
7. ルーティングを追加 (`app/routes/routes.go`)
8. Base構造体を更新
9. Swaggerドキュメントを生成 (`make swagger`)

### テスト

```bash
# 全テストを実行
make test

# 詳細モードで実行
make test-v

# 特定のパッケージのみ
docker-compose run --rm app go test ./app/usecase/... -v
```

### コード品質

```bash
# Lintチェック
make lint

# コードフォーマット
make fmt

# go vetチェック
make vet
```

## データベース

### スキーマ

主要なテーブル：
- `users` - ユーザー情報
- `products` - 商品マスタ（バーコード管理）
- `tenants` - 店舗マスタ
- `inventories` - 在庫情報（商品×店舗）
- `price_settings` - 価格設定（期間・セール対応）
- `orders` - 注文情報
- `order_items` - 注文明細

すべてのテーブルで論理削除（soft delete）をサポートしています。

### マイグレーション

```bash
# マイグレーション実行
make migrate

# DBシェルに接続
make db-shell

# テーブル確認
\dt
```

### シードデータ

アプリケーション起動時に自動的にシードデータが投入されます。
シードデータはupsert方式で、重複を防ぎます。

## トラブルシューティング

### ポートが使用中の場合

```bash
# 8080番ポートの使用状況確認
lsof -i :8080

# docker-compose.yml でポートを変更
```

### DBのリセット

```bash
# 完全リセット
make reset

# より強力なクリーンアップ
make clean-all
make reset
```

### ログの確認

```bash
# すべてのログ
make logs

# アプリケーションのログのみ
make logs-app
```

## 環境変数

```bash
# データベース設定（docker-compose.yml内）
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=myapp

# Ginフレームワーク設定
GIN_MODE=debug  # 本番環境では "release" に設定
```

## ライセンス

MIT License

## 貢献

プルリクエストを歓迎します。大きな変更の場合は、まずissueを作成して変更内容を議論してください。

1. プロジェクトをフォーク
2. フィーチャーブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をコミット (`git commit -m 'Add some amazing feature'`)
4. ブランチにプッシュ (`git push origin feature/amazing-feature`)
5. プルリクエストを作成

## 作者

Your Name - [@yourtwitter](https://twitter.com/yourtwitter)

Project Link: [https://github.com/yourusername/gin-practice](https://github.com/yourusername/gin-practice)
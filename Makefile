.PHONY: help build up down start stop restart logs ps status clean clean-all reset migrate seed test lint fmt vet tidy shell db-shell swagger

# 変数定義
DOCKER_COMPOSE := docker-compose
APP_NAME := gin-practice
DB_NAME := myapp
DB_USER := user

# デフォルトターゲット
.DEFAULT_GOAL := help

# ヘルプ
help:
	@echo "Gin Practice - Makefile Commands"
	@echo "================================"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make build        - Dockerイメージをビルド"
	@echo "  make up           - コンテナをバックグラウンドで起動"
	@echo "  make down         - コンテナを停止して削除"
	@echo "  make start        - 既存のコンテナを開始"
	@echo "  make stop         - コンテナを停止（削除しない）"
	@echo "  make restart      - コンテナを再起動"
	@echo "  make logs         - コンテナのログを表示（フォロー）"
	@echo "  make logs-app     - アプリケーションのログのみ表示"
	@echo "  make ps           - コンテナの状態を表示"
	@echo "  make status       - 詳細なステータスを表示"
	@echo ""
	@echo "Database Commands:"
	@echo "  make migrate      - データベースマイグレーションを実行"
	@echo "  make seed         - シードデータを投入"
	@echo "  make reset        - DBをリセットして再起動（データ削除）"
	@echo "  make db-shell     - PostgreSQLシェルに接続"
	@echo ""
	@echo "Development Commands:"
	@echo "  make run          - アプリケーションを起動（docker-compose up）"
	@echo "  make dev          - 開発モード（ログ表示付き）で起動"
	@echo "  make test         - テストを実行"
	@echo "  make test-v       - テストを詳細モードで実行"
	@echo "  make lint         - golangci-lintでコードをチェック"
	@echo "  make fmt          - コードをフォーマット"
	@echo "  make vet          - go vetでコードをチェック"
	@echo "  make tidy         - go.modを整理"
	@echo "  make shell        - アプリケーションコンテナのシェルに接続"
	@echo ""
	@echo "API Documentation:"
	@echo "  make swagger      - Swagger APIドキュメントを生成"
	@echo "  make swagger-fmt  - Swaggerコメントをフォーマット"
	@echo ""
	@echo "Cleanup Commands:"
	@echo "  make clean        - コンテナとボリュームを削除"
	@echo "  make clean-all    - すべてのDockerリソースを削除"
	@echo "  make prune        - 未使用のDockerリソースを削除"

## Docker Commands ##

# Dockerイメージをビルド
build:
	$(DOCKER_COMPOSE) build --no-cache

# コンテナをバックグラウンドで起動
up:
	$(DOCKER_COMPOSE) up -d
	@echo "アプリケーションが起動しました: http://localhost:8080"
	@echo "Swagger UI: http://localhost:8080/swagger/index.html"

# コンテナを停止して削除
down:
	$(DOCKER_COMPOSE) down

# 既存のコンテナを開始
start:
	$(DOCKER_COMPOSE) start

# コンテナを停止（削除しない）
stop:
	$(DOCKER_COMPOSE) stop

# コンテナを再起動
restart:
	$(DOCKER_COMPOSE) restart

# ログを表示（フォロー）
logs:
	$(DOCKER_COMPOSE) logs -f

# アプリケーションのログのみ表示
logs-app:
	$(DOCKER_COMPOSE) logs -f app

# コンテナの状態を表示
ps:
	$(DOCKER_COMPOSE) ps

# 詳細なステータスを表示
status: ps
	@echo ""
	@echo "Docker Volumes:"
	@docker volume ls | grep $(APP_NAME) || echo "No volumes found"

## Database Commands ##

# データベースマイグレーションを実行
migrate:
	$(DOCKER_COMPOSE) run --rm migrate
	@echo "マイグレーションが完了しました"

# シードデータを投入
seed:
	@echo "シードデータを投入中..."
	$(DOCKER_COMPOSE) exec app go run main.go
	@echo "シードデータの投入が完了しました"

# DBをリセットして再起動
reset: clean up migrate
	@echo "データベースがリセットされました"

# PostgreSQLシェルに接続
db-shell:
	$(DOCKER_COMPOSE) exec db psql -U $(DB_USER) -d $(DB_NAME)

## Development Commands ##

# アプリケーションを起動（フォアグラウンド）
run:
	$(DOCKER_COMPOSE) up

# 開発モードで起動（ビルド＋ログ表示）
dev: build
	$(DOCKER_COMPOSE) up

# テストを実行
test:
	$(DOCKER_COMPOSE) run --rm app go test ./...

# テストを詳細モードで実行
test-v:
	$(DOCKER_COMPOSE) run --rm app go test -v ./...

# Lintを実行
lint:
	$(DOCKER_COMPOSE) run --rm app golangci-lint run

# コードをフォーマット
fmt:
	$(DOCKER_COMPOSE) run --rm app go fmt ./...

# go vetでチェック
vet:
	$(DOCKER_COMPOSE) run --rm app go vet ./...

# go.modを整理
tidy:
	$(DOCKER_COMPOSE) run --rm app go mod tidy

# アプリケーションコンテナのシェルに接続
shell:
	$(DOCKER_COMPOSE) exec app /bin/sh

## API Documentation ##

# Swagger APIドキュメントを生成
swagger:
	$(DOCKER_COMPOSE) run --rm app swag init -g main.go -o docs
	@echo "Swagger documentation generated at http://localhost:8080/swagger/index.html"

# Swaggerコメントをフォーマット
swagger-fmt:
	$(DOCKER_COMPOSE) run --rm app swag fmt

## Cleanup Commands ##

# コンテナとボリュームを削除
clean:
	$(DOCKER_COMPOSE) down --volumes --remove-orphans
	docker volume ls -q | grep -E "($(APP_NAME)|postgres_data)" | xargs -r docker volume rm 2>/dev/null || true
	@echo "コンテナとボリュームを削除しました"

# すべてのDockerリソースを削除
clean-all: clean
	docker system prune -a -f --volumes
	@echo "すべてのDockerリソースを削除しました"

# 未使用のDockerリソースを削除
prune:
	docker system prune -f
	@echo "未使用のDockerリソースを削除しました"
# run in squad-admin-panel

BACKEND_PATH = backend

include .env

prod:
	docker-compose -f docker-compose.prod.yml up --build -d


dev-compose:
	docker-compose -f docker-compose.yml up --build -d

dev-apps:
	@cd ${BACKEND_PATH} && \
	BACKEND_DATABASE_URL="host=localhost user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=5432 sslmode=disable" \
	BACKEND_PORT=${BACKEND_PORT} \
	GIN_MODE=${GIN_MODE} \
	go run ./cmd/server/main.go

sqlc-gen:
	@cd ${BACKEND_PATH} && sqlc generate

dbmate-up:
	dbmate -u "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/main?sslmode=disable" up

dbmate-down:
	dbmate -u "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/main?sslmode=disable" down

.PHONY: prod dev-compose dev-apps sqlc-gen dbmate-up dbmate-down
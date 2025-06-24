# run in squad-admin-panel

FRONTEND_PATH = apps/frontend
BACKEND_PATH = apps/backend
REUSE_ENV = BACKEND_DATABASE_URL="host=localhost user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} port=5432 sslmode=disable" \
						KAFKA_URL=${KAFKA_URL} \
						KAFKA_TOPIC=${KAFKA_TOPIC} \
						KAFKA_EXTERN_URL=${KAFKA_EXTERN_URL} \

include .env

prod:
	docker-compose -f docker-compose.prod.yml up --build -d


dev-compose:
	docker-compose -f docker-compose.yml up --build -d

dev-apps:
	@echo "Starting backend and frontend in parallel..."
	@trap 'kill %1 %2' EXIT; \
	(cd ${BACKEND_PATH} && \
	 ${REUSE_ENV} \
	 BACKEND_PORT=${BACKEND_PORT} \
	 GIN_MODE=${GIN_MODE} \
	 go run ./cmd/server/main.go) & \
	(cd ${FRONTEND_PATH} && yarn && BACKEND_PORT=${BACKEND_PORT} FRONTEND_PORT=${FRONTEND_PORT} yarn dev) & \
	wait

kafka-script:
	@cd	${BACKEND_PATH} && \
	${REUSE_ENV} \
	go run ./cmd/kafka/main.go

sqlc-gen:
	@cd ${BACKEND_PATH} && sqlc generate

dbmate-up:
	dbmate -u "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/main?sslmode=disable" up

dbmate-down:
	dbmate -u "postgres://${DB_USER}:${DB_PASSWORD}@localhost:${DB_PORT}/main?sslmode=disable" down

.PHONY: prod dev-compose dev-apps dev-kafka sqlc-gen dbmate-up dbmate-down
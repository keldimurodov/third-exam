CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh

swag-gen:
	swag init -g ./api/router.go -o api/docs

run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "postgresql://postgres:123@localhost:5432/exam?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgres://postgres:123@localhost:5432/exam?sslmode=disable" -verbose down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq exam

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:123@localhost:5432/exam?sslmode=disable" force 1

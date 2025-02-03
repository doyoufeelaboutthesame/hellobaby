DB_DSN := "postgres://postgres:yourpassword@localhost:5432/main?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run main.go

gen:
	powershell -Command "oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml | Out-File -Encoding UTF8 ./internal/web/tasks/api.gen.go"
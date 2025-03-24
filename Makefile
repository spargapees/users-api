
migrateup:
	migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/users-api-db?sslmode=disable" -verbose up


migratedown:
	migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/users-api-db?sslmode=disable" -verbose down

dev:
	go run main.go
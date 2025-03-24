
migrateup:
	migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable" -verbose up


migratedown:
	migrate -path migrations -database "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable" -verbose down

dev:
	go run main.go
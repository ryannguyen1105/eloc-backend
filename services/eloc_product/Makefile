DB_Source_Local ?= "postgresql://root:secret123@localhost:5432/eloc_product?sslmode=disable"

migrateup:
	migrate -path db/migration -database $(DB_Source_Local) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_Source_Local) -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
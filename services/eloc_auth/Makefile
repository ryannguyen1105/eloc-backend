DB_Source_Local ?= "postgresql://root:secret123@localhost:5432/eloc_auth?sslmode=disable"

migrateup:
	migrate -path db/migration -database $(DB_Source_Local) -verbose up

migratedown:
	migrate -path db/migration -database $(DB_Source_Local) -verbose down
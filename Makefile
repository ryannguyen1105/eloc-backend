postgres:
	docker run --name eloc-postgres \
  --network eloc-network \
  -p 5432:5432 \
  -e POSTGRES_USER=root \
  -e POSTGRES_PASSWORD=secret123 \
  -v eloc-pgdata:/var/lib/postgresql/data \
  -d postgres:18-alpine

createdb_auth:
	docker exec -it eloc-postgres createdb --username=root --owner=root eloc_auth

createdb_product:
	docker exec -it eloc-postgres createdb --username=root --owner=root eloc_product

createdb_order:
	docker exec -it eloc-postgres createdb --username=root --owner=root eloc_order

createdb_all:	createdb_auth	createdb_product		createdb_order

dropdb_auth:
	docker exec -it eloc-postgres dropdb -f eloc_auth

dropdb_product:	
	docker exec -it eloc-postgres dropdb -f eloc_product

dropdb_order:
	docker exec -it eloc-postgres dropdb -f eloc_order

dropdb_all:	dropdb_auth	dropdb_product	dropdb_order



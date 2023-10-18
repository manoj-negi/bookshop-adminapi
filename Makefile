network:
	docker network create bank-network

createdb:
	docker exec -it postgres createdb --username=root --owner=root bookselling

dropdb:
	docker exec -it postgres dropdb simple_bank

migratecommand:
	migrate create -ext sql -dir db/migrations -seq  <table_name>

migrateup:
	migrate -path db/migrations -database "$(DB_URI)" -verbose up

migrateup1:
	migrate -path db/migrations -database "$(DB_URI)" -verbose up 1

migratedown:
	migrate -path db/migrations -database "$(DB_URI)" -verbose down

migratedown1:
	migrate -path db/migrations -database "$(DB_URI)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techschool/simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/techschool/simplebank/worker TaskDistributor

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis

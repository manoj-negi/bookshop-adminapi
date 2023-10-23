rename test.env to .en
task new_schema name=users //for create new migration file
task migrateup

go build -o ./bookshop-admin cmd/api/main.go
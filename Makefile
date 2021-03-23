cassandra:
	docker run --name cassandra3 -p 7000:7000 -p 7001:7001 -p 7199:7199 -p 9160:9160 -p 9042:9042 -d cassandra:3

redis:
	docker run --name redis6 -p 6379:6379 -d redis:6

startredis:
	docker start redis6

stopredis:
	docker stop redis6

startdb:
	docker start cassandra3

stopdb:
	docker stop cassandra3

migratedb:
	docker cp ./db/migration/create_db.cql cassandra3:/
	docker exec -d cassandra3 cqlsh localhost -f /create_db.cql

dropdb:
	docker cp ./db/migration/drop_db.cql cassandra3:/
	docker exec -d cassandra3 cqlsh localhost -f /drop_db.cql

dev:
	go run main.go

.PHONY: cassandra migratedb dropdb redis startredis stopredis

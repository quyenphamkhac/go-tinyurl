cassandra:
	docker run --name cassandra3 -p 9042:9042 -d cassandra:3

migratedb:
	docker cp ./db/migration/create_db.cql cassandra3:/
	docker exec -d cassandra3 cqlsh localhost -f /create_db.cql

dropdb:
	docker cp ./db/migration/drop_db.cql cassandra3:/
	docker exec -d cassandra3 cqlsh localhost -f /drop_db.cql

.PHONY: cassandra migratedb dropdb

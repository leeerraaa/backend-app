postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres
	
createdb:
	docker exec -it postgres-container createdb --username=root --owner=root crud-languages

dropdb:
	docker exec -it postgres-container dropdb crud-languages

migrateup:
	migrate -path ./schema -database "postgresql://root:root@localhost:5432/crud-languages?sslmode=disable" -verbose up

migratedown:
	migrate -path ./schema -database "postgresql://root:root@localhost:5432/crud-languages?sslmode=disable" -verbose down

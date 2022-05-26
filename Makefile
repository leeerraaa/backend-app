postgres:
	docker run --name postgres-container -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=qwerty -d postgres
	
createdb:
	docker exec -it postgres-container createdb --username=postgres --owner=postgres documents

dropdb:
	docker exec -it postgres-container dropdb documents

migrateup:
	migrate -path ./schema -database "postgresql://postgres:qwerty@localhost:5432/documents?sslmode=disable" -verbose up

migratedown:
	migrate -path ./schema -database "postgresql://postgres:qwerty@localhost:5432/documents?sslmode=disable" -verbose down

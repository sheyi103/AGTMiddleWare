mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest 

createdb:
	docker exec -it mysql mysql -uroot -psecret -e "create database agt_middleware_db;"

migrateup:
	migrate -path db/migration  -database "mysql://agt:Password123@tcp(localhost:3306)/agt_middleware_db" -verbose up

migratedown:
	migrate -path db/migration  -database "mysql://agt:Password123@tcp(localhost:3306)/agt_middleware_db" -verbose down
dropdb:
	docker exec -it mysql mysql -uroot -psecret -e "drop database agt_middleware_db;"

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: mysql createdb dropdb migrateup migratedown sqlc test
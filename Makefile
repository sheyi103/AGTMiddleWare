mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=@bbcSip! -d mysql:latest 
pullmysql:
	docker pull mysql:latest

createdb:
	docker exec -it mysql mysql -uroot -p@bbcSip! -e "create database agt_middleware_db;"

migrateup:
	migrate -path db/migration  -database "mysql://root:@bbcSip!@tcp(localhost:13306)/agt_middleware_db" -verbose up

migratedown:
	migrate -path db/migration  -database "mysql://root:@bbcSip!@tcp(localhost:13306)/agt_middleware_db" -verbose down
dropdb:
	docker exec -it mysql mysql -uroot -p@bbcSip! -e "drop database agt_middleware_db;"

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

.PHONY: mysql createdb dropdb migrateup migratedown sqlc test server


# mysql:
# 	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:latest 
# pullmysql:
# 	docker pull mysql:latest

# createdb:
# 	docker exec -it mysql mysql -uroot -psecret -e "create database agt_middleware_db;"

# migrateup:
# 	migrate -path db/migration  -database "mysql://root:secret@tcp(localhost:13306)/agt_middleware_db" -verbose up

# migratedown:
# 	migrate -path db/migration  -database "mysql://root:secret@tcp(localhost:13306)/agt_middleware_db" -verbose down
# dropdb:
# 	docker exec -it mysql mysql -uroot -psecret -e "drop database agt_middleware_db;"

# sqlc: 
# 	sqlc generate

# test: 
# 	go test -v -cover ./...

# server:
# 	go run main.go

# .PHONY: mysql createdb dropdb migrateup migratedown sqlc test server

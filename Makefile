build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down

mysql-client:
	docker-compose exec db mysql -uroot -proot -hlocalhost todo

migration:
	docker-compose exec migration bash

migration/up:
	docker-compose exec migration bash -c "sql-migrate up"

migration/down:
	docker-compose exec migration bash -c "sql-migrate down"

migration/status:
	docker-compose exec migration bash -c "sql-migrate status"

server/run:
	docker run -d --name gqlgen-todo -p 8080:8080 furiko/gqlgen-todo:latest

server/stop:
	docker container stop gqlgen-todo
	docker container rm gqlgen-todo

image/build:
	docker build -t furiko/gqlgen-todo:latest .
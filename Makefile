server/run:
	docker run -d --name gqlgen-todo -p 8080:8080 furiko/gqlgen-todo:latest

server/stop:
	docker container stop gqlgen-todo
	docker container rm gqlgen-todo

image/build:
	docker build -t furiko/gqlgen-todo:latest .
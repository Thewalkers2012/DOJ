.PHONY: createmysql createredis startmysql startredis

createmysql:
	sudo docker run --name DOJ_MYSQL -p 13306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

createredis:
	sudo docker run --name DOJ_REDIS -p 16379:6379 -d redis:6.2.5-alpine

startmysql:
	sudo docker start DOJ_MYSQL

startredis:
	sudo docker start DOJ_REDIS

test:
	go test -v -cover ./backend/...

server:
	go run ./backend/main.go

run:
	DB=go_todo DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 go run .
test:
	go test -v ./src/tests/...
gen:
	wire di/wire.go
db:
	docker-compose up -d
stop-db:
	docker-compose down

run:
	DB=go_todo DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 go run .
test:
	DB=go_todo_test DB_USER=root DB_PASSWORD=password DB_HOST=localhost DB_PORT=3306 go test -v ./...
gen:
	wire di/wire.go
db:
	docker-compose up -d
stop-db:
	docker-compose down
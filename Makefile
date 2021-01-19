compose-app:
	- docker-compose -f docker-compose-app.yml up --build --remove-orphans

compose-dbs:
	- docker-compose -f docker-compose-dbs.yml up --remove-orphans

gen:
	- @echo "generating..."
	- go generate -v ./graphql/...
	- @echo "done generating! ✅"

start:
	- go run --race server.go

build:
	- go build --race -o main .

migrate-up:
	- migrate -database postgres://dev:dev@localhost:5432/go_ms_1_dev?sslmode=disable -path db/migrations up

migrate-down:
	- migrate -database postgres://dev:dev@localhost:5432/go_ms_1_dev?sslmode=disable -path db/migrations down 

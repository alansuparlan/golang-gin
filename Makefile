PROJECT_NAME=alantestapp

# DATABASE
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=test_db
DB_SSL=disable

install:
	cd .. && go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate && \
	go get -u github.com/swaggo/swag/cmd/swag && go get -u github.com/cosmtrek/air && \
	go get github.com/vektra/mockery/v2/.../ && \
	cd ${PROJECT_NAME} && swag init

local:
	air -c config/.air.toml

test:
	go test -v -cover -coverprofile=cover.out ./...

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

dev-up:
	docker-compose -f docker-compose-dev.yml up -d --build

dev-down:
	docker-compose -f docker-compose-dev.yml down

prod-up:
	docker-compose up -d --build

prod-down:
	docker-compose down

migrate-up:
	migrate -source file:./script/migration/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} up

migrate-down:
	migrate -source file:./script/migration/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} down

migrate-drop:
	migrate -source file:./script/migration/ -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL} drop

scan:
	sonar-scanner

mockery:
	cd module/payment && mockery --name=Usecase && mockery --name=Repository
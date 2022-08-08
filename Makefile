BINARY=api_platforma

# GOLANG
run:
	go run ./src/cmd/main.go

watch:
	air

build:
	go build -o ./build/api_platforma ./src/cmd/main.go

lint:
	golangci-lint run ./... -v	

#  DOCKER
up-prod:
	echo "Starting docker prod environment"
	docker compose -f docker-compose.yml up --build

up-dev:
	echo "Starting docker Develop environment"
	docker compose -f docker-compose.dev.yml up -d --build

build-d:
	go build -o ${BINARY} ./src/cmd/main.go

.DEFAULT_GOAL: watch

###### DEV ######
start-dev:
	make up-dev && air

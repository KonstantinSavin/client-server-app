build:
		go build -v ./cmd

up:
	docker-compose up --force-recreate

down:
	docker-compose down
	docker image rm server-app-server

.DEFAULT_GOAL := build
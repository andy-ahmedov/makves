all: up

run:
	sudo docker-compose up server

up:
	docker-compose up -d postgresdb
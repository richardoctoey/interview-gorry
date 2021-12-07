all:
	build
build:
	docker build -t gorry .
stop:
	docker-compose stop
run:
	docker-compose up -d
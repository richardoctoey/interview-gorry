all:
	build
build:
	docker build -t gorry-richard-oey .
stop:
	docker-compose stop
run:
	docker-compose up -d
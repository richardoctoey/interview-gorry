all:
	build
build:
	docker build -t xapiens .
stop:
	docker-compose stop
run:
	docker-compose up -d
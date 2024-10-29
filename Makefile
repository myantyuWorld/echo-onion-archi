up:
	docker-compose up
down:
	docker-compose down
build:
	docker-compose up --build
api:
	docker-compose exec api /bin/bash
ps:
	docker ps

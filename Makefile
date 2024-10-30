up:
	docker-compose up
down:
	docker-compose down
build:
	docker-compose up --build
api:
	docker-compose exec api /bin/bash

# psql -d book_loan -U postgres
# \z // テーブル一覧表示
# \d TABLE_NAME テーブル定義確認
db:
	docker-compose exec db /bin/bash
ps:
	docker ps

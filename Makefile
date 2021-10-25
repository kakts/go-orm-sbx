.PHONY: up
up: # MySQLのDockerコンテナをdocker-composeにより起動
	docker-compose up -d

# MySQLのDockerコンテナをdocker-composeにより停止
.PHONY: down
down:
	docker-compose down

# MySQL初期化スクリプトの実行
.PHONY: init
init:
	./init-mysql.sh


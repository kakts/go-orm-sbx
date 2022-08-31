.PHONY: up
up: # MySQLのDockerコンテナをdocker-composeにより起動
	docker compose up -d

# MySQLのDockerコンテナをdocker-composeにより停止
.PHONY: down
down:
	docker compose down

.PHONY: logs
logs:
	docker compose logs

# MySQL初期化スクリプトの実行
.PHONY: init
init:
	./init-mysql.sh

.PHONY: mysql
mysql:
	mysql -udocker -pdocker -h "127.0.0.1"

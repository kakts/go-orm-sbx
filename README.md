# go-orm-sbx
go orm用のsbx

# Setting up a MySQL container

## MySQLのDockerコンテナ起動

```
make up
```
or
```
docker-compose up -d
```

## MySQL初期化スクリプトの実行
```
make init
```
or
```
./init-mysql.sh
```
これにより、MySQLコンテナ内でMySQLの初期化スクリプトを実行する


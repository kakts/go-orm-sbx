# gorm
gorm用app

# how to run
## run gorm app
```
make run
```


## ローカル開発の場合
起動時にMySQLへ接続するため、ローカル環境にMySQLを立ち上げている必要があります。
このリポジトリのトップディレクトリにdocker-compose.ymlで、設定しており、docker composeで関連するコンポーネントを起動してから、 gorm appを起動してください

```
$ pwd
/Users/hoge/fuga/go-orm-sbx

$ make up
# MySQLなどを起動

$ cd gorm_app
$ make run

> [GIN-debug] Listening and serving HTTP on :8080

起動後、8080ポートへリクエストをして疎通確認

$ curl localhost:8080
{"message":"Hello world"}%  
```




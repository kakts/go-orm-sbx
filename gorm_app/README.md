# gorm
gorm用app

# How to run
## run gorm app
```
make run
```

## API Document
- Access to localhost:8080/swagger/index.html

# How to develop

## start up MySQL
Start up MySQL container from Make command located in root directry.
```
cd ../; make up
```

## setting up swagger tools
- install swag.
```
make install
```

- check if `swag` command can be used.
```
swag

NAME:
   swag - Automatically generate RESTful API documentation with Swagger 2.0 for Go.
...
```

- Update API document.
```
$ make doc
swag init
2022/03/09 01:50:28 Generate swagger docs....
2022/03/09 01:50:28 Generate general API Info, search dir:./
2022/03/09 01:50:28 create docs.go at  docs/docs.go
2022/03/09 01:50:28 create swagger.json at  docs/swagger.json
2022/03/09 01:50:28 create swagger.yaml at  docs/swagger.yaml

```


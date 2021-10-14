[![Go Report Card](https://goreportcard.com/badge/github.com/3crabs/card-collection)](https://goreportcard.com/report/github.com/3crabs/card-collection)
# card-collection
Платформа для коллекционирования карт
## Configurations
Configuration used gonfig. You have to create 'dev_config.json'. Example:
```json
{
  "DB_USERNAME": "usename",
  "DB_PASSWORD": "pass",
  "DB_PORT": "5432",
  "DB_HOST": "localhost",
  "DB_NAME": "name"
}
```

## 🏃‍♀️ Run
```sh
git clone https://github.com/3crabs/card-collection
// run
go run main.go
// build
go build main.go
./main.go
```

## 📌 Feature
- REST API server with [Echo Framework](https://github.com/labstack/echo)
- Dataebase ORM using [GORM](https://github.com/jinzhu/gorm)
- JWT Authentication(Not yet)
- Echo Custome Validator [go-playground/validator](https://github.com/go-playground/validator)

## 📚 Project Stack
- echo
- gorm (Postgres)
- go-playground/validator
- gonfig
- godotenv(Not yet)
- jwt-go(Not yet)

## Architecture
| Folder | Details |
| --- | ---|
| controller | Holds the api endpoints |
| storage | Database Initializer and DB manager |
| route | router setup |
| models | Models|

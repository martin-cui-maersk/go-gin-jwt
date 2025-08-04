# go-gin-jwt

用GIN框架编写的JWT示例



## 配置文件

```sh
cp env.example .env
```



## 目录结构

```sh
├── controllers
│   └── auth.go
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── middlewares
│   └── middlewares.go
├── models
│   ├── mysql.go
│   └── user.go
├── README.md
├── routes
│   └── route.go
└── utils
    ├── global
    │   └── result.go
    └── token
        └── token.go
```



## 运行

```sh
go run main.go
```



## 接口示例

### POST   /api/auth/register        --> go-gin-jwt/controllers.Register (4 handlers)

```sh
# curl
curl --location 'localhost:8000/api/auth/register' \
--header 'Content-Type: application/json' \
--data '{
    "username":"abc",
    "password":"123"
}'

# response
{
    "code": 200,
    "msg": "register success",
    "data": {
        "username": "abc",
        "password": "123"
    }
}
```

### POST   /api/auth/login           --> go-gin-jwt/controllers.Login (4 handlers)

```sh
# curl
curl --location 'localhost:8000/api/auth/login' \
--header 'Content-Type: application/json' \
--data '{
    "username":"abc",
    "password":"123"
}'

# response
{
    "code": 200,
    "msg": "Success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQzMDI5MzksImlhdCI6MTc1NDI5NTczOSwiaXNzIjoiand0IGV4YW1wbGUiLCJuYmYiOjE3NTQyOTU3MzksInN1YiI6Inh4eC5hYmMuY29tIiwidXNlcl9pZCI6MX0.Rf2gYluJm3USyPeA_Zg0LxTLfXtnELOcSvb1rTINrfQ"
    }
}
```

### GET    /api/user/info            --> go-gin-jwt/controllers.CurrentUser (5 handlers)

```sh
# curl
curl --location --request GET 'localhost:8000/api/user/info' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQzMDI5MzksImlhdCI6MTc1NDI5NTczOSwiaXNzIjoiand0IGV4YW1wbGUiLCJuYmYiOjE3NTQyOTU3MzksInN1YiI6Inh4eC5hYmMuY29tIiwidXNlcl9pZCI6MX0.Rf2gYluJm3USyPeA_Zg0LxTLfXtnELOcSvb1rTINrfQ' \
--data '{
    "username":"abc",
    "password":"123"
}'

# response
{
    "code": 200,
    "msg": "Success",
    "data": {
        "ID": 1,
        "CreatedAt": "2025-08-04T16:18:55+08:00",
        "UpdatedAt": "2025-08-04T16:18:55+08:00",
        "DeletedAt": null,
        "username": "abc",
        "password": ""
    }
}

```




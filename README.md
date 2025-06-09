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


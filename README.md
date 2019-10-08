# 接口文档

## 文件结构

```
.
├── client
│   ├── setup.go
│   ├── status.go
│   ├── test.go
│   └── t.go
├── main
│   ├── auth.go
│   ├── database.go
│   ├── main.go
│   ├── server.go
│   ├── test.db
│   └── web.go
├── README.md
└── web
    ├── certificate
    ├── css
    ├── img
    ├── index.html
    ├── js
    ├── login.html
    ├── part
    ├── site
    ├── status
    └── webfonts

```

## 接口

### 状态码

- 000 ERROR
- 200 OK
- 201 OK and MORE
- 404 Not Found

### 身份认证

- jwt
- token for each server

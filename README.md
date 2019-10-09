# XvA-Go

## 文件结构

```
.
├── README.md
├── src
│   ├── Client
│   │   ├── config.json
│   │   ├── data.json
│   │   └── setup.go
│   └── Server
│       ├── auth.go
│       ├── Certificate.go
│       ├── database.go
│       ├── Domain.go
│       ├── main.go
│       ├── run.sh
│       ├── ServerApi.go
│       ├── Server.go
│       ├── Site.go
│       ├── test.db
│       └── User.go
└── web
    ├── certificate
    │   ├── add.html
    │   ├── index.html
    │   └── info.html
    ├── css
    │   ├── all.min.css
    │   ├── bootstrap.min.css
    │   └── custom.css
    ├── img
    │   └── logo.PNG
    ├── index.html
    ├── js
    │   ├── bootstrap.min.js
    │   ├── checkPower.js
    │   ├── gconfig.js
    │   └── jquery-3.4.1.min.js
    ├── login.html
    ├── part
    │   ├── footer.html
    │   └── nav.html
    ├── site
    │   ├── add.html
    │   ├── manage.html
    │   ├── site.conf
    │   └── siteInfo.html
    ├── status
    │   ├── domain.html
    │   ├── server.html
    │   └── serverInfo.html
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

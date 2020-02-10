# ControlCenter

## 文件

```
src/
    Client - 子客户端
    Server - 母服务端
```

## API 接口

```
POST    /user/auth/login
POST    /user/auth/register


POST    /server/setup/:token
POST    /update/:token POST
GET     /server/update/:token
GET     /server/Certificate/:token/:id

GET     /

POST    /system/restart

- - -
auth 2.0 request
- - -
POST    /web/debug/check
GET     /web/ServerInfo
PUT     /web/ServerInfo
GET     /web/ServerInfo/Certificate
DELETE  /web/Server/:ip
GET     /web/DomainInfo
PUT     /web/DomainInfo
PUT     /web/UserInfo/:mail/:key/:value
PUT     /web/SiteInfo
GET     /web/SiteInfo
PUT     /web/Certificate
GET     /web/Certificate
POST    /web/Certificate
DELETE  /web/Certificate
PUT     /web/link/Certificate/:ServerID/:CerID
DELETE  /web/link/Certificate/:ServerID/:CerID
POST    /web/backup
GET     /web/:mail/:pass/backup
GET     /web/UserInfo/Password/:oldpass/:newpass
GET     /web/UserInfo/Token
PUT     /web/UserInfo/Token
GET     /web/UserInfo
```

## Install 安装

### Client

```
./Client install https://127.0.0.1:1323 `hostname` `curl ip.sb -4` `curl ip.sb -6` user_token
```
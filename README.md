# 接口文档

## 文件结构

```
XvA-Go
    |--server
    |   |--server.go
    |   |--database.go
    |   |--auth.go
    |
    |--client
    |
    |--web
    |   |--index.html
```

## 接口

### 状态码

- 000 ERROR
- 200 OK
- 201 OK and MORE
- 404 Not Found

### 身份认证

- jwt
- token for each client

### WEB

| URL | 参数 | 示例 | 返回 |
| --- | --- | --- | --- |
| auth/login |(POST) username,password||(JSON) {status:(200/000),jwt:(jwt/none)}|
| auth/register | (POST) username,password,mail| |(JSON) {status:(200/000)}|
| web/setup | (GET) ipv4,*ipv6(option)*,hostname| |(JSON) {status:(200/000),token:(TOKEN)}|
| web/update | (GET) id,token,info(json2string) ||(JSON) {status:(200/000),info:(none/errorinfo)}|
| web/get | (GET) id,token | |(JSON) {status:(200/201/000),info:(JSON)}|
| server/getUpdate | (POST) jwt | | (JSON) |
| server/getFile | (POST) jwt,path | | (String) |
| server/editFile | (POST) jwt,(File) | |(JSON) {status:(200/000)} |


### server.go

### database.go

def getData(table string,)

### auth.go

## 数据库结构

### Users

| | |
| --- | --- |
| id | int |
| mail | string |
| username | string |
| password | string |
| level | int |

## Servers

| | | |
| --- | --- | --- |
| id | int | |
| ipv4 | string | |
| ipv6 | string | (option) |
| hostname | string | |
| token | string | (only) |

```go
type User struct {
	ID       int64  `gorm:"AUTO_INCREMENT"`
	Username string `json:"name" xml:"name" form:"name" query:"name"`
	Mail     string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
	Level    int64
}

type Server struct {
	Hostname string
	Ipv4     string
	Ipv6     string
	ID       int64 `gorm:"AUTO_INCREMENT"`
}

type Service struct {
	ID   int64 `gorm:"AUTO_INCREMENT"`
	Name string
}

type Domain struct {
	ID     int64 `gorm:"AUTO_INCREMENT"`
	Name   string
	Status string
	Cer    string
	Key    string
}

type Config struct {
	ID          int64 `gorm:"AUTO_INCREMENT"`
	configKey   string
	configValue string
}

type Callback struct {
	Code int64
	Info string
}
```
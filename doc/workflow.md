# API 定义

## 约定

返回成功一般情况均为

```json
{
    "Code": 200,
    "Info": "OK"
}
```

## POST /user/auth/login 登录

请求
```json
{
    "email": "i@test.com",
    "password": "2333333"
}
```
返回-成功
```json
{
    "token": "eyJhbGciO............................JWT token.....................Mqy99QxMt2Snqts"
}
```
返回-失败
```json
{
    "Code": 0,
    "Info": "account or password incorrect"
}
```

## POST /user/auth/register 注册

请求
```json
{
    "name": "alex",
    "email": "i@test.com",
    "password": "1265234234"
}
```
返回-失败
```json
{
    "Code": 0,
    "Info": "This account has been used"
}
```

## POST /server/setup/"User Token" 服务器注册

请求
```json
{
    "hostname": "Ali-SZ",
    "ipv4": "8.8.8.8",
    "ipv6": "::0"
}
```
返回-成功
```json
{
    "Code": 200,
    "Info": "d8.........Server token...........0ebd"
}
```
返回-失败
```json
{
    "Code": 0,
    "Info": "Server already exists"
}
```

## GET /server/update/"Server token" 服务器获取控制信息

返回-成功
```json
{
  "Code": 200,
  "Sites": [],
  "Certificates": [
    {
      "ID": 1,
      "Domain": "test.com",
      "FullChain": "",
      "Key": ""
    }
  ],
  "Services": null,
  "Dockers": null
}
```
返回-失败
```json
{
  "Code": 0,
  "Info": "Unauthorized"
}
```
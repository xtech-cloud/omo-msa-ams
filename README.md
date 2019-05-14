# 简介

账户管理服务

支持http和rpc两种调用方式

# 编译

进入omo-msa-ams目录，执行compile.sh脚本


# 部署

设置环境变量

```
export GIN_MODE=release
export AMS_HTTP_ADDR=:80
export AMS_RPC_ADDR=:10080
export AMS_LOG_FILE=/var/log/ams.log
export AMS_LOG_LEVEL=INFO
```

- SQLite

设置以下环境变量

```
export AMS_DATABASE_DRIVER=sqlite
export AMS_SQLITE_FILEPATH=<db文件所在路径>
```

- Mysql

设置以下环境变量

```
export AMS_DATABASE_DRIVER=mysql
export AMS_MYSQL_ADDR=127.0.0.1:3306
export AMS_MYSQL_USER=<mysyql用户>
export AMS_MYSQL_PASSWORD=<mysql密码>
export AMS_MYSQL_DATABASE=ams
```

设置完环境变量后启动ams


# HTTP API

## `/ams/signup`

**简要描述:**

- 注册账号
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|username|是  |string |用户名|
|password|是  |string |密码|


**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
        "uuid":"0bc97a7ba9645b0a556803bf9e671002"
    }
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| uuid| string | 账号UUID|

**备注**

`ams直接存储收到的密码，不进行额外加密处理`


## `/ams/signin`

**简要描述:**

- 登录
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|username|是  |string |用户名|
|password|是  |string |密码|


**返回示例**

```json
{
    "code":200,
    "expire":"2019-05-14T15:50:13+08:00",
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTc4MjAyMTMsImlkIjoiNjg1ZDgxZWNjMDNlYzg0NzRmNGMyZWRhZmJhOGQzYzgiLCJvcmlnX2lhdCI6MTU1NzgxNjYxM30.JggbFwGDjIqo8UnvuUqOXEHnEg_Z6SJGrGXo7lg2V3s"
}
```

**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| code| int| http状态码 （200代表正常）|
| expire| string| token有效期|
| token| string| jwt token|



## `/ams/auth/signout`

**简要描述:**

- 登出
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |


**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
    }
}
```

**备注**

使用jwt头的ID，不需要传递额外参数

## `/ams/auth/current`

**简要描述:**

- 获取账号信息
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |


**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
        "profile":""
    }
}
```
**返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|:-----|
| profile | string | 账号简介|


**备注**

使用jwt头的ID，不需要传递额外参数

## `/ams/auth/update`

**简要描述:**

- 更新账户介绍
  
**请求方式：**

- POST 

**参数：** 

|参数名|必选|类型|说明|
|:----    |:---|:--- |:---   |
|profile|是|string|简介|


**返回示例**

```json
{
    "code":0, 
    "message":"",
    "data":{
    }
}
```

# SamBlog

一个使用golang开发的blog项目


### 安装

```shell
go get -u github.com/zituocn/samblog
```

### 编译运行

需要有golang的编译环境

```shell
cd samblog
go build 
./samblog
```

### 访问

```shell
http://127.0.0.1:8001
```

### 配置文件 

*conf/app.conf*

```shell
app_name = samblog  # 应用名
run_mode = dev      # 运行模式 dev && prod
auto_render = true  # 是否渲染html
http_addr = 8001    # 运行端口号
session_on = false  # 是否启用session
views = views       # html模板目录 
template_left = <<  # 模板符号
template_right = >> # 模板符号


defaultDB = samblog  # 数据库名 

[samblog]
user = "root"
password = "root"
host = "127.0.0.1"
port = 6606


[system]
xxKey = 9758a4f12f7d2f01  # 加密key
cookieAgeHour = 1         # cookie 过期时间(小时)
```
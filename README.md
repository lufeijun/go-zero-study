# go-zero-learn

参考：https://github.com/nivin-studio/go-zero-mall

# 一、 user 微服务

## 启动

```
# 启动 rpc 
cd service/user/rpc
go run user.go -f etc/user.yaml

# 启动 api
cd service/user/api
go run user.go -f etc/user.yaml
```

## 访问

```
http://127.0.0.1:8000/api/user/login

```

## 说明

### 启动多个 rpc

在启动一个 rpc 后，如果把 port 改为 9001 ，login 错误提示稍微修改 ```用户不存在111``` 。这样就启动了两个 rpc 服务器，即 127.0.0.1:9000、127.0.0.1:9001，
在访问 ```http://127.0.0.1:8000/api/user/login``` 时，设置错误的 uid。多次请求时，返回值为：```用户不存在、用户不存在111```。可以看出来，底层做了负载均衡





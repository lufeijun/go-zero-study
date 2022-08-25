# go-zero-learn


# user 微服务

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

#### websocketApi SDK

> 用户连接websocketApi的SDK,使用此SDK部署到生产环境需要做好`充分`的测试.  
> 本例仅供参考,部分逻辑需要根据自己的业务去实现  
> 例子:`handle/test.go`

#### 启动
> go run main.go -h api.aex.zone -p /v3 -s wss
- `-h` 主机名
- `-p` 路径
- `-s` wss or ws

> 目录结构
```
ws-api-cli
    ├── func                公共的函数
    ├── handle              业务逻辑
    │   └── test.go         例子
    ├── lib                 核心库文件
    │   ├── client.go
    │   ├── conn.go
    │   ├── interface.go    接口
    │   ├── read.go
    │   ├── warning.md
    │   └── write.go
    ├── types               公共类型文件
    │   ├── comm.go
    │   ├── define.go
    │   ├── req.go
    │   └── resp.go
    └── main.go             主函数

```

###### 业务逻辑请在`handle/work()`方法里面写,并实现`lib/Handle接口`

> 主要的功能:
- msg <- b.Receive 接收msg
- b.send <- msg 发送msg
- 如此简单,敬请享用吧,祝您早日暴富 :)

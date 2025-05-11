# mcp-ip
使用离线IP库获取IP对应的地理位置信息。

# ip库支持
当前使用的ip库为：https://github.com/lionsoul2014/ip2region。

# 使用
## 编译
```bash
$ git clone https://github.com/qiangmzsx/mcp-ip.git
$ cd mcp-ip
$ go build -o mcp-ip
```

## 启动参数
1. -transport 指定传输方式，默认为streamable_http，还也可以有stdio、see；
2. -state_mode 指定状态模式，默认为stateful，还可以有stateless；
3. -xdb_path 指定ip库的路径，默认为./data/ip2region.xdb；
4. -port 指定端口，默认为8080；

## stdio
```json
{
  "mcpServers": {
    "ipinfo": {
      "command": "mcp-ip",
      "args": [
        "-transport",
        "stdio",
        "-xdb_path",
        "./data/ip2region.xdb"
      ]
    }
  }
}
```
## streamable_http
```bash
http://127.0.0.1:8080/mcp
```


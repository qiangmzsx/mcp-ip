# mcp-ip
使用离线IP库获取IP对应的地理位置信息。

# ip库支持
当前使用的ip库为：https://github.com/lionsoul2014/ip2region。

# 使用
## 编译
```bash
$ git clone https://github.com/mcp-io/mcp-ip.git
$ cd mcp-ip
$ go build -o mcp-ip
```

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



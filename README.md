
# mcp-ip
Get geographic location information from IP addresses using offline IP database.

# IP Database Support
Currently using ip2region database: https://github.com/lionsoul2014/ip2region.

# Usage
## Build
```bash
$ git clone https://github.com/qiangmzsx/mcp-ip.git
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

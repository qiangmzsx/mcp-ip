
# mcp-ip
Get geographic location information from IP addresses using offline IP database.

# Supported IP Database
Currently using ip2region: https://github.com/lionsoul2014/ip2region.

# Usage
## Build
```bash
$ git clone https://github.com/qiangmzsx/mcp-ip.git
$ cd mcp-ip
$ go build -o mcp-ip
```

## Startup Parameters
1. -transport specifies transport method, default is streamable_http, options include stdio, see;
2. -state_mode specifies state mode, default is stateful, option is stateless;
3. -xdb_path specifies path to IP database, default is ./data/ip2region.xdb;
4. -port specifies port number, default is 8080;

## stdio Example
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
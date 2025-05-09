package ip

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var ip2RegionService *IP2RegionService
var onceForIp2RegionService = sync.Once{}

type IP2RegionService struct {
	path     string
	Searcher *xdb.Searcher
}

func NewIP2RegionService(xdbPath string) (service *IP2RegionService, errReturn error) {
	onceForIp2RegionService.Do(func() {
		buff := []byte{}
		_, err := os.Lstat(xdbPath)
		if !os.IsNotExist(err) {
			buff, err = os.ReadFile(xdbPath)
			if err != nil {
				errReturn = err
				return
			}
		}
		searcher, err := xdb.NewWithBuffer(buff)
		if err != nil {
			errReturn = err
			return
		}
		ip2RegionService = &IP2RegionService{
			path:     xdbPath,
			Searcher: searcher,
		}
	})

	return service, errReturn
}

func GetIP2Region(_ context.Context, request *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
	req := new(IPReq)
	if err := protocol.VerifyAndUnmarshal(request.RawArguments, &req); err != nil {
		return nil, err
	}
	// 校验IP地址格式
	ip := net.ParseIP(req.IP)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address format: %s", req.IP)
	}
	rawData, err := ip2RegionService.Searcher.SearchByStr(req.IP)
	if err != nil {
		return nil, err
	}
	text := XDB2Location(rawData).String()
	return &protocol.CallToolResult{
		Content: []protocol.Content{
			&protocol.TextContent{
				Type: "text",
				Text: text,
			},
		},
	}, nil
}

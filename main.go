package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/qiangmzsx/mcp-ip/ip"
)

func main() {
	var err error
	err = registerIPService()
	if err != nil {
		log.Fatalf("Failed to register IP service: %v", err)
		return
	}
	srv, err := server.NewServer(
		getTransport(),
		server.WithServerInfo(protocol.Implementation{
			Name:    "ip-server",
			Version: "1.0.0",
		}),
	)

	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}

	ip2regionTool, err := protocol.NewTool("ip2region", "Supports IPv4/IPv6 address retrieval for country, province, city, and IP information.", ip.IPReq{})
	if err != nil {
		log.Fatalf("Failed to create tool: %v", err)
	}

	srv.RegisterTool(ip2regionTool, ip.GetIP2Region)

	errCh := make(chan error)
	go func() {
		errCh <- srv.Run()
	}()

	if err = signalWaiter(errCh); err != nil {
		log.Fatalf("signal waiter: %v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
}

func registerIPService() error {
	var xdbPath string
	flag.StringVar(&xdbPath, "xdb_path", "./data/ip2region.xdb", "ip2region xdb path")
	// flag.Parse()

	_, err := ip.NewIP2RegionService(xdbPath)
	return err
}

func getTransport() (t transport.ServerTransport) {
	var mode, port, stateMode string
	flag.StringVar(&mode, "transport", "streamable_http", "The transport to use, should be \"stdio\" or \"sse\" or \"streamable_http\"")
	flag.StringVar(&port, "port", "8080", "sse server address")
	flag.StringVar(&stateMode, "state_mode", "stateful", "streamable_http server state mode, should be \"stateless\" or \"stateful\"")

	flag.Parse()

	switch mode {
	case "stdio":
		log.Println("start IP mcp server with stdio transport")
		t = transport.NewStdioServerTransport()
	case "sse":
		addr := fmt.Sprintf("127.0.0.1:%s", port)
		log.Printf("start IP mcp server with sse transport, listen %s", addr)
		t, _ = transport.NewSSEServerTransport(addr)
	case "streamable_http":
		addr := fmt.Sprintf("127.0.0.1:%s", port)
		log.Printf("start IP mcp server with streamable_http transport, listen %s", addr)
		t = transport.NewStreamableHTTPServerTransport(addr, transport.WithStreamableHTTPServerTransportOptionStateMode(transport.StateMode(stateMode)))
	default:
		panic(fmt.Errorf("unknown mode: %s", mode))
	}

	return t
}

func signalWaiter(errCh chan error) error {
	signalToNotify := []os.Signal{syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM}
	if signal.Ignored(syscall.SIGHUP) {
		signalToNotify = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, signalToNotify...)

	select {
	case sig := <-signals:
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			log.Printf("Received signal: %s\n", sig)
			// graceful shutdown
			return nil
		}
	case err := <-errCh:
		return err
	}

	return nil
}

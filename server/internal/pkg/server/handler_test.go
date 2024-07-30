package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/xxDENxx/powtcp_go/server/internal/pkg/config"
)

type InMemoryCache struct {
	dataMap map[string]string
}

func InitInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		dataMap: make(map[string]string, 0),
	}
}

func (c *InMemoryCache) Add(key string) error {
	c.dataMap[key] = "value"
	return nil
}

func (c *InMemoryCache) Get(key string) (bool, error) {
	value := c.dataMap[key]
	return value != "", nil
}

func (c *InMemoryCache) Delete(key string) {
	delete(c.dataMap, key)
}

func Test_handleRequest(t *testing.T) {
	t.Run("request challenge", func(t *testing.T) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "config", &config.Config{})
		ctx = context.WithValue(ctx, "cache", InitInMemoryCache())
		
		go func() { // func call
			l, _ := net.Listen("tcp", ":3000")
			defer l.Close()
			conn, _ := l.Accept()
			
			handleRequest(ctx, conn)
		}()
		
		time.Sleep(time.Second)
		conn, _ := net.Dial("tcp", ":3000")
		defer conn.Close()
		fmt.Fprintf(conn, "Challenge")
		conn.(*net.TCPConn).CloseWrite()
			
		buf, _ := io.ReadAll(conn)
		if len(buf) == 0 {
			t.Errorf("handleRequest() empty response")
		}
	})
	t.Run("request quote", func(t *testing.T) {
		cache := InitInMemoryCache()
		ctx := context.Background()
		ctx = context.WithValue(ctx, "config", &config.Config{})
		ctx = context.WithValue(ctx, "cache", cache)
		
		go func() { // func call
			l, _ := net.Listen("tcp", ":3000")
			defer l.Close()
			conn, _ := l.Accept()
			
			handleRequest(ctx, conn)
		}()
		
		time.Sleep(time.Second)
		challenge, _ := NewChallenge(ctx)
		challenge.Counter = 1000
		msg := fmt.Sprintf("%d:%d:%d", challenge.Seed, challenge.Zeros, challenge.Counter)
		conn, _ := net.Dial("tcp", ":3000")
		defer conn.Close()
		fmt.Fprint(conn, msg)
		conn.(*net.TCPConn).CloseWrite()

		buf, _ := io.ReadAll(conn)
		if len(buf) == 0 {
			t.Errorf("handleRequest() empty response")
		}
	})
}

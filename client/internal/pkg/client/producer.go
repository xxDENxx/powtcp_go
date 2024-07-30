package client

import (
	"context"
	"io"
	"log"
	"net"
	"time"
)



func Run(ctx context.Context) error {
	config := ctx.Value("config").(*Config)
	for {
		quote, err := receiveQuote(ctx, config)
		if err != nil {
			return err
		}
		log.Println("New quote:", quote)

		log.Println("Sleep for ", config.RequestIntervalDur())
		time.Sleep(config.RequestIntervalDur())
	}
}

func receiveQuote(ctx context.Context, config *Config) (string, error) {
	// Make first request to receive challenge
	resp, err := makeRequest(ctx, config, []byte("Challenge"))
	if err != nil {
		return "", err
	}
	challenge, err := makeChallenge(resp)
	if err != nil {
		return "", err
	}
	challenge.Calculate()
	// Make second request to receive quote
	resp, err = makeRequest(ctx, config, []byte(challenge.ToMsg()))
	if err != nil {
		return "", err
	}
	return resp, nil
}

func makeRequest(_ context.Context, config *Config, msg []byte) (string, error) {
	conn, err := net.Dial("tcp", config.ServerAddress())
	if err != nil {
		return "", err
	}
	defer conn.Close()
	
	_, err = conn.Write(msg)
	if err != nil {
		return "", err
	}
	
	conn.(*net.TCPConn).CloseWrite()
	
	buf, err := io.ReadAll(conn)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
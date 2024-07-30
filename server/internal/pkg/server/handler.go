package server

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/xxDENxx/powtcp_go/server/internal/pkg/config"
	"github.com/xxDENxx/powtcp_go/server/internal/pkg/repository"
)

func Run(ctx context.Context) error {
	conf := ctx.Value("config").(*config.Config)
	listener, err := net.Listen("tcp", conf.Address())
	if err != nil {
		return err
	}
	defer listener.Close()
	
	log.Println("Ready for connections")
	log.Println("listening", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("CONNECTION: ", err)
		}
		log.Println("connect received")

		go handleRequest(ctx, conn)
	}
}

func handleRequest(ctx context.Context, conn net.Conn) {
	defer conn.Close()
	buf, err := io.ReadAll(conn)
	if err != nil {
		log.Println("bad connection")
		return
	}
	if string(buf) == "Challenge" {
		challenge, err := NewChallenge(ctx)
		if err != nil {
			log.Println("Cache problem: ", err)
			return
		}
		conn.Write([]byte(challenge.ToMsg()))
	} else {
		challenge, err := ChallengeFromMsq(ctx, string(buf))
		if err != nil || !challenge.isValid() {
			log.Println("Bad challenge ", err)
			return
		}
		repo := repository.NewQuoteRepo()
		conn.Write([]byte(repo.GetQuote()))
	}
	conn.(*net.TCPConn).CloseWrite()
	log.Println("connection done")
}
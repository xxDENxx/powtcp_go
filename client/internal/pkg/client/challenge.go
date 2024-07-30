package client

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Challenge struct {
	Seed    int
	Zeros   int
	Counter int
	hash    []byte
}

func makeChallenge(raw string) (*Challenge, error) {
	parts := strings.Split(raw, ":")
	if len(parts) != 2 {
		return nil, errors.New("WRONG MESSAGE")
	}
	seed, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	zeros, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	if zeros > 40 {
		return nil, errors.New("TOO DIFFICULT")
	}

	return &Challenge{seed,zeros,0,[]byte("")}, nil
}

func (c *Challenge) ToMsg() string {
	return fmt.Sprintf("%d:%d:%d", c.Seed, c.Zeros, c.Counter)
}

func (c *Challenge) Calculate() {
	for {
		data := fmt.Sprintf("%d:%d:%d", c.Seed, c.Zeros, c.Counter)
		c.hash = getHash(data)
		if c.isCorrect() {
			log.Println("Calculated in ", c.Counter)
			return
		}
		c.Counter++
	}	
}

func (c *Challenge) isCorrect() bool {
	for _, ch := range c.hash[:c.Zeros] {
		if ch != byte('0') {
			return false
		}
	}
	return true
}

func getHash(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}

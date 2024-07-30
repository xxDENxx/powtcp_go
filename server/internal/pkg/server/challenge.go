package server

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/xxDENxx/powtcp_go/server/internal/pkg/config"
)

type Challenge struct {
	Seed    int
	Zeros   int
	Counter int
	ctx     context.Context
}

type Cache interface {
	// Add - add rand value with expiration (in seconds) to cache
	Add(string) error
	// Get - check existence of int key in cache
	Get(string) (bool, error)
	// Delete - delete key from cache
	Delete(string)
}

func NewChallenge(ctx context.Context) (*Challenge, error) {
	cache := ctx.Value("cache").(Cache)
	conf := ctx.Value("config").(*config.Config)
	challenge := &Challenge{0,conf.Difficult,0,ctx}
	err := challenge.generateSeed(cache)
	if err != nil {
		return nil, err
	}
	return challenge, nil
}

func ChallengeFromMsq(ctx context.Context, msg string) (*Challenge, error) {
	conf := ctx.Value("config").(*config.Config)
	parts := strings.Split(msg, ":")
	if len(parts) != 3 {
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
	if zeros != conf.Difficult {
		return nil, errors.New("WRONG DIFFICULT")
	}
	counter, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	return &Challenge{seed,zeros,counter, ctx}, nil
}

func (c *Challenge) generateSeed(cache Cache) error {
	randValue := 0
	for i := 1; i <= 5; i++ {
		randValue = rand.Intn(100000000)
		err := cache.Add(strconv.Itoa(randValue))
		if i == 5 && err != nil {
			return err
		}
	}
	c.Seed = randValue
	return nil
}

func (c *Challenge) ToMsg() string {
	return fmt.Sprintf("%d:%d", c.Seed, c.Zeros)
}

func (c *Challenge) isValid() bool {
	cache := c.ctx.Value("cache").(Cache)
	exist, err := cache.Get(strconv.Itoa(c.Seed))
	if err != nil || !exist {
		return false
	}
	hash := []byte(c.getHash())
	for _, ch := range hash[:c.Zeros] {
		if ch != byte('0') {
			return false
		}
	}
	// every challenge only single use
	cache.Delete(strconv.Itoa(c.Seed))
	return true
}

func (c *Challenge) getHash() string {
	data := fmt.Sprintf("%d:%d:%d", c.Seed, c.Zeros, c.Counter)
	h := sha256.New()
	h.Write([]byte(data))
	return string(h.Sum(nil))
}
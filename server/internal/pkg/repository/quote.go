package repository

import "math/rand"

type quoteRepo struct {
	quotes []string
}

// https://www.xavier.edu/jesuitresource/online-resources/quote-archive1/knowledge-quotes
func NewQuoteRepo() *quoteRepo {
	return &quoteRepo{
		[]string{
			"When the going gets rough - turn to wonder.",
			"If you have knowledge, let others light their candles in it.",
			"A bird doesn't sing because it has an answer, it sings because it has a song.",
			"We are not what we know but what we are willing to learn.",
			"Good people are good because they've come to wisdom through failure.",
		},
	}
}
func (repo *quoteRepo) GetQuote() string {
	return repo.quotes[rand.Intn(len(repo.quotes))]
}
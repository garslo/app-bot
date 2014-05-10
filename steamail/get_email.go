package steamail

import (
	"errors"
	"log"
	"net/mail"
	"time"

	"github.com/garslo/email"
)

var (
	ErrTooManyFetchRetries = errors.New("max fetch retries reached")
)

type steamCodeGetter struct {
	fetcher       email.EmailFetcher
	codeExtractor SteamCodeExtractor
	retryInterval time.Duration
	numRetries    int
}

func NewSteamCodeGetter(fetcher email.EmailFetcher) *steamCodeGetter {
	return NewSteamCodeGetter2(fetcher, 30*time.Second, 10, NewSteamCodeExtractor())
}

func NewSteamCodeGetter2(
	fetcher email.EmailFetcher,
	retryInterval time.Duration,
	numRetries int,
	codeExtractor SteamCodeExtractor,
) *steamCodeGetter {
	return &steamCodeGetter{
		fetcher:       fetcher,
		retryInterval: retryInterval,
		numRetries:    numRetries,
		codeExtractor: codeExtractor,
	}
}

func NewGmailSteamCodeGetter(username, password string) SteamCodeGetter {
	fetcher := email.NewGmailFetcher(username, password)
	return NewSteamCodeGetter(fetcher)
}

func (g *steamCodeGetter) GetSteamCode() (string, error) {
	log.Println("Fetching login code")
	for i := 0; i < g.numRetries; i++ {
		msgs, err := g.fetcher.FetchEmails()
		if err == nil {
			code, err := g.getCode(msgs)
			if err == nil {
				return code, nil
			}
		}
		log.Println("Could not fetch code, retrying in %s")
		time.Sleep(g.retryInterval)
	}
	return "", ErrTooManyFetchRetries
}

func (g *steamCodeGetter) getCode(msgs []*mail.Message) (string, error) {
	var code string
	var err error
	for _, msg := range msgs {
		code, err = g.codeExtractor.ExtractCode(msg)
		if err == nil {
			return code, nil
		}
	}
	return "", err
}

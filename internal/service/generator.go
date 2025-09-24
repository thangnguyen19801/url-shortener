package service

import (
	"github.com/teris-io/shortid"
	"github.com/yourusername/url-shortener/internal/model"
	"github.com/yourusername/url-shortener/internal/storage"
	"strings"
	"time"
)

type Generator interface {
	GenerateCode(url string, host string) (string, string, error)
}

type generator struct {
	store *storage.Postgres
}

func NewGenerator(store *storage.Postgres) Generator {
	return &generator{store: store}
}

func (g *generator) GenerateCode(url string, host string) (string, string, error) {
	code, err := shortid.Generate()
	if err != nil {
		return "", "", err
	}

	existing, err := g.store.FindByTarget(url)
	if err == nil && existing != nil {
		shortURL := buildShortURL(host, existing.Code)
		return existing.Code, shortURL, nil
	}

	u := &model.URL{
		Code:      code,
		Target:    url,
		CreatedAt: time.Now(),
	}
	if err := g.store.CreateURL(u); err != nil {
		return "", "", err
	}

	shortURL := buildShortURL(host, code)
	return code, shortURL, nil
}

func buildShortURL(host string, code string) string {
	scheme := "https://"

	host = strings.TrimRight(host, "/")
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return host + "/" + code
	}
	return scheme + host + "/" + code
}

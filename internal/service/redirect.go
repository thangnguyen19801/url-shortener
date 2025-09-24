package service

import (
	"github.com/yourusername/url-shortener/internal/storage"
	"log"
)

type Redirect interface {
	UpdateClickEvent(code string, referer string, ua string, ip string) (string, error)
}
type redirect struct {
	store *storage.Postgres
}

func NewRedirect(store *storage.Postgres) Redirect {
	return &redirect{store: store}
}

func (r *redirect) UpdateClickEvent(code string, referer string, ua string, ip string) (string, error) {
	u, err := r.store.GetByCode(code)
	if err != nil {
		return "", err
	}
	go func() {
		err := r.store.IncClicks(u, referer, ua, ip)
		if err != nil {
			log.Println(err)
		}
	}()
	return u.Target, nil
}

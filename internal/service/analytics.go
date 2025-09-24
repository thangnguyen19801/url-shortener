package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/url-shortener/internal/storage"
)

type Analytics interface {
	GetAnalytics(code string) (interface{}, error)
}
type analytics struct {
	store *storage.Postgres
}

func NewAnalytics(store *storage.Postgres) Analytics {
	return &analytics{store}
}

func (a *analytics) GetAnalytics(code string) (interface{}, error) {
	u, err := a.store.GetAnalytics(code)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"code":       u.Code,
		"target":     u.Target,
		"clicks":     u.Clicks,
		"created_at": u.CreatedAt,
		"last_seen":  u.LastSeen,
	}, nil
}

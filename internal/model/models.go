package model

import "time"

type URL struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"uniqueIndex;size:16" json:"code"`
	Target    string    `gorm:"not null" json:"target"`
	CreatedAt time.Time `json:"created_at"`
	Clicks    uint64    `gorm:"default:0" json:"clicks"`
	LastSeen  *time.Time `json:"last_seen"`
}

type ClickEvent struct {
	ID        uint      `gorm:"primaryKey"`
	URLID     uint      `gorm:"index"`
	Referer   string
	UA        string
	IP        string
	CreatedAt time.Time
}

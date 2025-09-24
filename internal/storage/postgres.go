package storage

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/yourusername/url-shortener/internal/model"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	gcfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(postgres.Open(dsn), gcfg)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.URL{}, &model.ClickEvent{}); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(50)
		sqlDB.SetConnMaxLifetime(time.Minute * 5)
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) CreateURL(u *model.URL) error {
	return p.db.Create(u).Error
}

func (p *Postgres) GetByCode(code string) (*model.URL, error) {
	var u model.URL
	if err := p.db.Where("code = ?", code).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (p *Postgres) IncClicks(u *model.URL, referer, ua, ip string) error {
	now := time.Now()
	if err := p.db.Model(u).Updates(map[string]interface{}{"clicks": gorm.Expr("clicks + 1"), "last_seen": &now}).Error; err != nil {
		return err
	}
	ce := model.ClickEvent{URLID: u.ID, Referer: referer, UA: ua, IP: ip, CreatedAt: now}
	return p.db.Create(&ce).Error
}

func (p *Postgres) GetAnalytics(code string) (*model.URL, error) {
	return p.GetByCode(code)
}

func (p *Postgres) FindByTarget(target string) (*model.URL, error) {
	var url model.URL
	if err := p.db.Where("target = ?", target).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

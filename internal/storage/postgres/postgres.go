package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"link-shortener/internal/cfg"
	"link-shortener/internal/model"
	"link-shortener/internal/storage"
	"link-shortener/internal/utils"
)

type PostgresInstance struct {
	db *gorm.DB
}

func NewPostgresInstance(repo *PostgresInstance) storage.StorageInstance {
	return repo
}

func (p *PostgresInstance) Setup() {
	dbCfg := cfg.DbConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.SSLMode)
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = p.db.AutoMigrate(&model.Shortening{})
	if err != nil {
		log.Println(err)
	}
}

// TODO:Добавить проверку на уникальность
func (p *PostgresInstance) Save(longUrl string, urlLen int) (string, error) {
	shortUrl := utils.Hash_func(longUrl, urlLen)
	shorty := model.Shortening{
		OriginalURL: longUrl,
		ShortUrl:    shortUrl,
	}
	err := p.db.Create(&shorty).Error
	if err != nil {
		log.Println(err)
	}
	return shortUrl, err
}

func (p *PostgresInstance) Unique(shortUrl string, longUrl string) bool {
	var shorty model.Shortening
	r := p.db.Where("short_url = ?", shortUrl).Find(&shorty)
	return r.RowsAffected == 0
}

func (p *PostgresInstance) Get(shortUrl string) (string, error) {
	var shorty model.Shortening
	tx := p.db.Where("short_url = ?", shortUrl).First(&shorty)
	if tx.Error != nil {
		return "", tx.Error
	}
	return shorty.OriginalURL, nil
}
package postgres

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"link-shortener/internal/cfg"
	"link-shortener/internal/model"
	"link-shortener/internal/storage"
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
	//Если надо очистить таблицу перед работой с ней
	if dbCfg.Format {
		stmt := &gorm.Statement{DB: p.db}
		stmt.Parse(&model.Shortening{})
		tableName := stmt.Schema.Table
		del := fmt.Sprintf("DELETE FROM %s", tableName)
		p.db.Exec(del)
	}
}

func (p *PostgresInstance) Save(longUrl string, shortUrl string) error {
	shorty := model.Shortening{
		OriginalURL: longUrl,
		ShortUrl:    shortUrl,
	}
	tx := p.db.Begin()
	err := tx.Create(&shorty).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	return tx.Commit().Error
}

func (p *PostgresInstance) Unique(shortUrl string, longUrl string) bool {
	var shorty model.Shortening
	p.db.Where("short_url = ?", shortUrl).Find(&shorty)
	return (shorty.ShortUrl == "") || (longUrl == shorty.OriginalURL)
}

func (p *PostgresInstance) Get(shortUrl string) (string, error) {
	var shorty model.Shortening
	tx := p.db.Begin()
	req := tx.Where("short_url = ?", shortUrl).First(&shorty)
	if req.Error != nil {
		tx.Rollback()
		return "", req.Error
	}
	return shorty.OriginalURL, tx.Commit().Error
}

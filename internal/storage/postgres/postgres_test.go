package postgres

import (
	"fmt"
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"link-shortener/internal/cfg"
	"link-shortener/internal/model"
)

const path = "../../cfg"

func testSetup() *PostgresInstance {
	var p = new(PostgresInstance)
	cfg.Init(path)
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
	return p
}

func TestGet(t *testing.T) {
	p := testSetup()
	//p.db.Exec("DELETE FROM shorties")
	shortUrl := "someShortUrl"
	longUrl := "someUrl"
	err := p.db.Create(&model.Shortening{
		OriginalURL: longUrl,
		ShortUrl:    shortUrl,
	}).Error
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	res, err := p.Get(shortUrl)
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	assert.Equal(t, longUrl, res)
}

func TestSave(t *testing.T) {
	p := testSetup()
	shortUrl := "SELcdoeR9j"
	longUrl := "https://www.youtube.com/watch?v=GtL1huin9EE"
	res, err := p.Save(longUrl, 10)
	if err != nil {
		t.Errorf("error occured: %v", err)
	}
	assert.Equal(t, shortUrl, res)
}

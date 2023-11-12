package storage

// StorageInstance - интерфейс хранилища данных, задающий все нужные методы
type StorageInstance interface {
	// Unique проверяет уникальность сокращенного URL
	Unique(shortUrl string, longUrl string) bool
	// Save сохраняет оригинальный URL в хранилище данных и возвращает уникальный сокращенный URL
	Save(longUrl string, shortUrl string) error
	// Get возвращает оригинальный URL из хранилища данных по сокращенному URL
	Get(shortUrl string) (string, error)
}

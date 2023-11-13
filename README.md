# link-shortener
 Implementation of a service that provides an API for creating short links(using gRPC and gRPC Gateway)
- Docker
  - Создать docker образ сервиса с параметром database=in-memory/postgres (по умолчанию значение database=in-memory)
  ```bash
  make df database=postgres
  ```
  - Запуск сервиса и базы данных в Docker 
  ```bash
  make service_up
  ```
- Для запуска приложения без docker
  ```bash
  make run
  ```

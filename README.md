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
- Для запуска приложения локально надо поменять адрес HTTP на "localhost" в internal/cfg/config.yml
  ```bash
  go mod init link-shortener
  go mod tidy
  make run
  ```
*Примечание:программа запускалась на Ubuntu.В силу специфики системы(для создания docker требуются права суперпользователя) возможно потребуется запустить альтернативные команды при запуске на другом OS:
- Docker
  - Создать docker образ сервиса с параметром database=in-memory/postgres (по умолчанию значение database=in-memory)
  ```bash
  docker build --build-arg DATABASE=$(database) --tag link-shortener .
  ```
  - Запуск сервиса и базы данных в Docker 
  ```bash
  docker-compose -f docker-compose.yml up -d --remove-orphans
  ```

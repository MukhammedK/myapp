# MyApp

Веб-приложение на Go с разделением ролей пользователей: админ и обычный пользователь.

## Технологии

- Go 1.21
- PostgreSQL
- HTML Templates

## Возможности

- Подключение к PostgreSQL
- Разделение окружений dev/prod через `.env`
- Разные роли пользователей: admin и user

## Запуск

1. Создай файл `.env.dev`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=твой_пароль
DB_NAME=postgres
ENV=dev
ADMIN_NAME=Admin
ADMIN_ROLE=admin
```

2. Запусти сервер:
```bash
go run main.go
```

Сервер запустится на `http://localhost:8080`

# О проекте

## HTTP-server
Проект, представленный в данном репозитории, является HTTP-server'ом. 

### API
API описан в `internal\handlers\handler.go`, предоставляет CRUD-методы данные записываются в `"DTO"` после чего передаются в слой сервиса.

### Service
Методы Service'а описаны в интерфейсе `"Service"`, находящемся в  `internal\handlers\handler.go`, сами же методы
находятся в `internal\service`.

### Repository
Методы Repository описаны в интерфейсе `"Repo"`, находящемся в  `internal\service\service.go`, сами же методы
находятся в `internal\repo`.
Данные хранятся в Postgres. SQL запросы осуществляются с помощью [GORM](https://github.com/go-gorm/gorm).

#### Миграции
Миграции описаны в `migrations`.
Для применения миграций используется [goose](https://github.com/pressly/goose).

#### Тесты
Тесты находятся в соответствующих тестируемым слоям директориях, так тесты для сервисного слоя находятся в `internal\service\service_test.go`.
Для тестирования используется  [testify](https://github.com/stretchr/testify).

## Docker

Описан докерфайл для образа HTTP-server'а

## Makefile

В мейкфайле описаны команды для локального запуска приложения, применения миграций, сборки докер-образов, и их запуска.

## Запуск приложения

Запуск осуществляется через команду:
```docker-compose up```
 после чего на порту `:8080` приложение будет доступно для работы


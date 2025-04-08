[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)
[![docker](https://img.shields.io/badge/-Docker-464646?style=flat-square&logo=docker)](https://www.docker.com/)

# url-shortener
# Сервис, предоставляющий API по созданию сокращённых ссылок

---
## Технологии
* Go 1.23.5
* REST API
* Docker
* Postman

---
## Взаимодейстивие с сервисом
### Запуск сервера
```
docker-compose up-d
```


*Пример POST запроса на адрес* `http://localhost:8081/url`:

**Request:**
```JSON
{
    "url": "https://ya.ru"
}
```
**Response:**
```JSON
{
    "status": "OK",
    "alias": "asBS1"
}
```
*Пример GET запроса на адрес* `http://localhost:8000/asBS1`:

**Response:**
```JSON
{
    "status": "OK",
    "url": "https://ya.ru"
}
```
*Повторный POST запрос на адрес* `http://localhost:8000/url`:
**Response:**
```JSON
{
    "status": "Error",
    "error": "url already exists"
}
```
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)

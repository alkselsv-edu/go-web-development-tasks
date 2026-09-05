# Валидация HTTP-запросов

В этом упражнении мы реализуем обработчик `POST /users` для сервиса, который работает с пользователями из Азии.

Нужно принять данные пользователя, проверить их и сохранить в память. Для простоты используем мапу, где ключом является логин пользователя.

## app/solution.go

В запросе `POST /users` приходит JSON с полями:

* Логин пользователя (`username string`).
* Email (`email string`).
* Возраст (`age int`).
* Страна (`country string`).

Проверьте данные по правилам:

* `username` заполнен и содержит только латиницу в нижнем регистре и цифры.
* `email` заполнен и имеет корректный формат.
* `age` в диапазоне от 18 до 130.
* `country` принимает одно из значений: `"Japan"`, `"Vietnam"`, `"Malaysia"`, `"Thailand"`.

Если пользователь с таким `username` уже существует, его данные нужно перезаписать.

Если валидация не прошла, верните HTTP-статус `422 Unprocessable Entity` и текст ошибки. Для валидации используйте библиотеку [go-playground/validator](https://github.com/go-playground/validator).

## Примеры

Корректный запрос:

```bash
curl -X POST -H "Content-Type: application/json" http://localhost:8080/users -d '{"username": "john1", "email": "john@doe.com", "age": 25, "country": "Japan"}'
```

Ожидаемый ответ:

```bash
HTTP/1.1 200 OK
```

Некорректный запрос:

```bash
curl -X POST -H "Content-Type: application/json" http://localhost:8080/users -d '{"username": "John_1", "email": "invalid", "age": 2000, "country": "Unknown"}'
```

Пример ответа:

```bash
HTTP/1.1 422 Unprocessable Entity

Key: 'CreateUserRequest.Username' Error:Field validation for 'Username' failed on the 'username' tag
Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag
Key: 'CreateUserRequest.Age' Error:Field validation for 'Age' failed on the 'lte' tag
Key: 'CreateUserRequest.Country' Error:Field validation for 'Country' failed on the 'oneof' tag
```

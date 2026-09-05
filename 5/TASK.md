# Сериализация данных в JSON

Представим, что вы разрабатываете веб-приложение, которое находит индекс числа в массиве, используя [бинарный поиск](https://ru.hexlet.io/courses/basic-algorithms/lessons/binary-search/theory_unit).

## app/solution.go

Веб-приложение должно уметь обрабатывать запрос вида:

```bash
POST /search

{
  "numbers": [1,2,3,4,5],
  "target": 3
}
```

Веб-приложение должно рассчитать местоположение числа "target" в массиве "numbers". При этом массив чисел "Numbers" всегда отсортирован по возрастанию. На данный HTTP-запрос веб-приложение должно вернуть следующий HTTP-ответ:

```bash
HTTP/1.1 200 OK

{
  "target_index": 2
}
```

"target_index" — это индекс искомого числа "target" в массиве "numbers". Отсчет индексов идет с нуля. Если число не было найдено, то нужно вернуть HTTP-ответ с кодом *404 Not Found* и телом с индексом *-1* и ошибкой "Target was not found":

```bash
HTTP/1.1 404 Not Found

{
  "target_index": -1,
  "error": "Target was not found"
}
```

Также если в веб-приложение приходит HTTP-запрос с телом в некорректном формате JSON, то оно должно вернуть HTTP-ответ с кодом *400 Bad Request* и ошибку в теле с текстом "Invalid JSON":

```bash
HTTP/1.1 400 Bad Request

{
  "target_index": -1,
  "error": "Invalid JSON"
}
```

В стандартной Go-библиотеке `slices` есть функция `Index()`, которая выполняет поиск в массиве:

```go
targetIndex := slices.Index(numbers, target)
```

Если индекс не был найден, то функция возвращает -1.

## Подсказки

* [Index](https://pkg.go.dev/slices@go1.22.4#Index)

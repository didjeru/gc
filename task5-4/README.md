# Задание №5-4

## Условие задачи

Напишите утилиту для копирования файлов, используя пакет flag.

## Запуск

```bash
go run main.go
```

## Тестирование

```bash
$ go run main.go
Usage of /tmp/go-build092540240/b001/exe/main:
  -dst string
        destination file
  -src string
        source file
exit status 1
```

```bash
$ go run main.go -src 1.txt -dst 2.txt
copied "1.txt" to "2.txt"
```



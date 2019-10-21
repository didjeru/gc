# Задание №5-5

## Условие задачи

Напишите упрощенный аналог утилиты grep.

## Запуск

```bash
go build grep.go
```

## Тестирование

```bash
$ ls -la | ./grep
Usage: ./grep [-e] pattern
  -e string
        RegExp pattern
```

```bash
$ ls -la | ./grep go
-rw-r--r--   1 user  staff      714 17 окт 18:55 grep.go
```

```bash
$ ls -la | ./grep -e "[0-9]"
итого 2476
drwxrwxr-x  2 di di    4096 окт 17 19:29 .
drwxrwxr-x 21 di di    4096 окт 17 19:11 ..
-rwxrwxr-x  1 di di 2518171 окт 17 19:29 grep
-rw-r--r--  1 di di     714 окт 17 19:11 grep.go
-rw-r--r--  1 di di     793 окт 17 19:12 README.md
```

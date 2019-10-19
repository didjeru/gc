# Задание №4-3

## Условие задачи

Дописать функцию, которая будет выводить справку к калькулятору. Она должна вызываться при вводе слова **help** вместо выражения.

## Запуск

```bash
go run main.go
```

## Тестирование

```bash
Умный калькулятор. Для помощи, наберите help
Введите необходимые вычисления > cos(%)
Ошибка - 1:5: expected operand, found '%'
Умный калькулятор. Для помощи, наберите help
Введите необходимые вычисления > cos(5)
Результат: 0.2836621854632263
Умный калькулятор. Для помощи, наберите help
Введите необходимые вычисления > der
Ошибка - Неизвестная константа der
Умный калькулятор. Для помощи, наберите help
Введите необходимые вычисления > help
Помощь по калькулятору:
Функции: sqrt, abs, log, ln, sin, cos, tan, arcsin, arccos, arctan, max, min
Писать без проблеов, например: 5+5 или min(75.5,65.4) или cos(6)
Умный калькулятор. Для помощи, наберите help
Введите необходимые вычисления > exit
$
```
# Golang Conf 2023

"School Language" interpreter

demo.sl

```
А равно 10
Б равно 20

показать('Вычисляем Ц...')

Ц равно 3 умножить на А плюс Б умножить на 2

показать('Вычисляем Г...')

Г равно 15 плюс Ц

Г равно Г минус 3

показать('Результат:')
показать(Г)
```

output:

```
Вычисляем Ц...
Вычисляем Г...
Результат:
82
```

## Install

Install go tools with `goyacc`

```shell
go get -u golang.org/x/tools/...
```

## Run

```shell
go generate
go run .
```

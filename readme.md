# Golang Conf 2023

"School Language" interpreter

demo.sl

```
А равно 10
Б равно 20

Ц равно 3 умножить на А плюс Б умножить на 2

Г равно 15 плюс Ц
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

Output:

```
А = 10
Б = 20
Ц = 70
Г = 85
Ответ: Г = 85
```

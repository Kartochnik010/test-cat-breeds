# Test task - Cat Breeds 

## About
Это консольное приложение на языке Go предназначено для извлечения списка пород кошек через API [Cat Fact Ninja](https://catfact.ninja/breeds). Оно группирует названия пород по стране происхождения, сортирует их по длине названия и сохраняет результаты в файл `out.json`. В проекте уделено особое внимание качеству кода и обработке ошибок.

## Usage

1. Клонируйте репозиторий:
```Bash
   git clone https://github.com/Kartochnik010/test-cat-breeds.git
```
2. Перейдите в директорию проекта:
```Bash
    cd test-cat-breeds
```
3. Скачайте зависимости:
```Go
    go mod tidy
```

3. Запустите приложение:
```Go
    make run
```
## Project structure
```
.
├── cmd
│   └── cli
│       └── main.go 
├── go.mod
├── config
│   └── config.go   
├── Makefile
├── internal
│   ├── models
│   │   └── models.go
│   └── api
│       └── api.go
├── utils
│   ├── files.go
│   └── utils.go
├── go.sum
├── readme.md
└── out.json
```
`/cmd` - точка входа в приложение.

`/config` - конфиги для приложения

`/internal/api` - пакет, содержащий функционал для работы с API.

`/internal/models` - пакет, определяющий структуры данных.

`/util` - пакет с вспомогательными функциями для обработки данных.

### API
Приложение делает запросы к API [Cat Fact Ninja](https://catfact.ninja/breeds), получая данные о породах кошек. Данные обрабатываются для группировки по стране происхождения и сортировки по длине названия.

### Error handling
Все сетевые запросы и операции обработки данных сопровождаются проверками на наличие ошибок. В случае возникновения ошибок, приложение предоставляет информативное сообщение об ошибке и завершает свою работу.

### Data output
Результаты обработки данных записываются в файл out.json, который создается в корневой директории проекта.

### Contribute to project
## Todos
- Print stack trace when encountering an error for better debugging
- 3rd party logger?
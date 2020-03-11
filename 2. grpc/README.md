# Запуск

Получение .go сервиса по .proto файлу:

```sh
$ protoc -I hello/ hello/hello.proto --go_out=plugins=grpc:hello
```
Для компиляции и запуска требуется перейти в корень директории grpc(.../grpc/).

Сервер

```sh
$ go run server/server.go
```

Клиент

```sh
$ go run client/client.go
```

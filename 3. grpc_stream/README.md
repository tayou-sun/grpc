# Запуск

Получение .go сервиса по .proto файлу:

```sh
$ protoc -I translit/ translit/translit.proto --go_out=plugins=grpc:translit
```
Для компиляции и запуска требуется перейти в корень директории grpc_stream (.../grpc_stream/).

Сервер

```sh
$ cd server
$ go run *.go
```

Клиент

```sh
$ go run client/client.go
```

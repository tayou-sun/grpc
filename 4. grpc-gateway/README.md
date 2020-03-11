*http интерфейс к grpc сервису*

# Запуск

Получение .go сервисjd по .proto файлу:

## сгенерировать .pb.go
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I grpc-gateway \
  --go_out=plugins=grpc:. \
  session.proto

## gateway
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  session.proto

Для компиляции и запуска требуется перейти в корень директории route_guide (.../examples/route_guide/).

Сервер

```sh
$ cd server
$ go run *.go
```

## HTTP-запросы
* curl -X POST -k http://localhost:8080/v1/session/create -H "Content-Type: text/plain" -d '{"login":"login", "useragent": "chrome"}'

* curl http://localhost:8080/v1/session/check/XVlBzgbaiC
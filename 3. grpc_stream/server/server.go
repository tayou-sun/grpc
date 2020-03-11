package main

import (
	"coursera/microservices/grpc_stream/translit"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//устанавливается порт для прослушивания подключений по протоколу TCP
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalln("Произошла ошибка", err)
	}

	//создание grpc сервера
	server := grpc.NewServer()

	//регестрирование сервера
	translit.RegisterTransliterationServer(server, NewTr())

	fmt.Println("Сервер стартовал на порту :8081")

	//Запуск сервера
	server.Serve(lis)
}

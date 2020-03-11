package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"

	"coursera/microservices/grpc_stream/translit"
)

func main() {

	//установка соединения с сервером
	grcpConn, err := grpc.Dial(
		"127.0.0.1:8081",
		//обозначает, что байты не будут шифроваться
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Ошибка: невозможно соединиться с сервером")
	}
	//defer - функция grcpConn.Close() будет выполнена
	//перед завершением работы программы
	defer grcpConn.Close()

	//созданание клиента
	tr := translit.NewTransliterationClient(grcpConn)

	//создание контекста
	//Контекст - это просто сборник мета-данных, ассоциированных с каким-то запросом
	//возвращает пустой контекст
	ctx := context.Background()

	client, err := tr.EnRu(ctx)

	// WaitGroup - синхронизация горутин
	wg := &sync.WaitGroup{}

	//обьявление, что будет две горутины
	wg.Add(2)

	//отправка данных потока
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		words := []string{"privet", "kak", "dela"}
		for _, w := range words {
			fmt.Println("-> ", w)
			client.Send(&translit.Word{
				Word: w,
			})
			time.Sleep(time.Millisecond)
		}
		client.CloseSend()
	}(wg)

	//получение данных с потока
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			outWord, err := client.Recv()
			if err == io.EOF {
				fmt.Println("\tstream closed")
				return
			} else if err != nil {
				return
			}
			fmt.Println(" <-", outWord.Word)
		}
	}(wg)

	wg.Wait()

}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"coursera/microservices/gateway/session"
)

func main() {

	//запуск одноверменно двух серверов
	//прокси и grpc-сервера
	proxyAddr := ":8080"
	serviceAddr := "127.0.0.1:8081"

	go gRPCService(serviceAddr)
	HTTPProxy(proxyAddr, serviceAddr)
}

func gRPCService(serviceAddr string) {

	lis, err := net.Listen("tcp", serviceAddr)

	if err != nil {
		log.Fatalln("Ошибка", err)
	}

	server := grpc.NewServer()

	session.RegisterAuthCheckerServer(server, NewSessionManager())

	server.Serve(lis)
}

func HTTPProxy(proxyAddr, serviceAddr string) {
	grcpConn, err := grpc.Dial(
		serviceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Ошибка", err)
	}
	defer grcpConn.Close()

	//создание мультиплексора
	//сопоставление запроса и обрабатывающей функции
	grpcGWMux := runtime.NewServeMux()

	err = session.RegisterAuthCheckerHandler(
		context.Background(),
		grpcGWMux,
		grcpConn,
	)
	if err != nil {
		log.Fatalln("Ошибка", err)
	}

	mux := http.NewServeMux()

	// отправляем в прокси только то, что нужно
	mux.Handle("/v1/session/", grpcGWMux)

	mux.HandleFunc("/", helloWorld)

	log.Fatal(http.ListenAndServe(proxyAddr, mux))

}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL:", r.URL.String())
}

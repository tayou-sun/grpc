package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/route_guide/hello"
)

const (
	address = "localhost:50051"
)

func timingInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf(`--
	call=%v
	req=%#v
	reply=%#v
	time=%v
	err=%v
`, method, req, reply, time.Since(start), err)
	return err
}

func main() {
	//Для создания потока ввода через буфер применяется
	reader := bufio.NewReader(os.Stdin)
	// Установка клиентского соединения с сервером
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(timingInterceptor),
	) //, grpc.WithBlock())

	if err != nil {
		log.Fatalf("Не удалось установить соединение: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	for {
		text, _ := reader.ReadString('\n')

		name := text

		if len(os.Args) > 1 {
			name = os.Args[1]
		}

		ctx := context.Background()

		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})

		if err != nil {
			log.Fatalf("Ошибка! %v", err)
		}

		log.Printf("Log! : %s", r.GetMessage())
	}

}

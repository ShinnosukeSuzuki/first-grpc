package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	hellopb "first-grpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  hellopb.GreetingServiceClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// gRPCクライアントを生成
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Hello()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Hello() {
	fmt.Println("Please enter your name.")
	scanner.Scan()
	name := scanner.Text()

	// リクエストに使うHelloRequest型の生成
	req := &hellopb.HelloRequest{
		Name: name,
	}

	// Helloメソッドの実行 -> HelloResponse型のレスポンスresを入手
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		// resの内容を標準出力に出す
		fmt.Println(res.GetMessage())
	}
}

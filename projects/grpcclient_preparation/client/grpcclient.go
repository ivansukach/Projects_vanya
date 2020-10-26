package client

import (
	"context"
	"github.com/ivansukach/grpc-server/protocol"
	"google.golang.org/grpc"
	"log"
)

func main(){
	log.Println("Клиент запущен ...")
	opts := grpc.WithInsecure()
	clientConnIntrfc, err := grpc.Dial("localhost:1323", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer clientConnIntrfc.Close()
	client := protocol.NewGetResponseClient(clientConnIntrfc)
	request := &protocol.GRRequest{Req: "Ping"}


	response, _ := client.GiveResponse(context.Background(), request)//Выполняется в другом потоке.
	// Пакет context в go позволяет вам передавать данные в вашу программу в каком-то «контексте».
	// Пакет context дает нам вести обмен данными, и содержит все необходимое для этого обмена
	log.Println("Ответ сервера:", response.GetRes())

}

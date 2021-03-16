//Package definition
package main

//Imports
import (
	"context"
	"log"

	chat "grpcTutorial/chatserver"

	"google.golang.org/grpc"
)

//Main definition
func main() {

	//We create a new Client and we assign credentials using  grpc.Dial
	//We assing the port where the server is listening, in this case port 9000
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No conecté: %s", err)
	}
	defer conn.Close()

	//Stub creation
	c := chat.NewChatServiceClient(conn)

	//Service method call, we send a MessageRequest and receive in response a MessageReply
	response, err := c.SendMessage(context.Background(), &chat.MessageRequest{Name: "gRPC Client"})
	if err != nil {
		log.Fatalf("Error al llamar a SendMessage: %s", err)
	}
	log.Printf("Respuesta del servidor: %s", response.Body)
}
  
//Package chat definition
package chat

//Imports
import (
	"context"
	"log"
)

//Server struct definition
type Server struct {
}

//SendMessage function implementation
func (s *Server) SendMessage(ctx context.Context, in *MessageRequest) (*MessageReply, error) {
	log.Printf("Mensaje recibido de : %s", in.GetName())
	return &MessageReply{Body: "Hola " + in.GetName() + " desde el servidor!"}, nil
}
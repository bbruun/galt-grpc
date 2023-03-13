package client

import (
	"log"
	// minion "github.com/bbruun/galt/proto"
)

// type myMinionClient struct {
// 	minion.UnimplementedMinionServiceServer
// }

// func (m *myMinionClient) CommandFromMinion(ctx context.Context, input *minion.CommandToMinion) (*minion.CommandFromMinion, error) {

// 	// a := time.Now().UTC()
// 	return &minion.CommandFromMinion{
// 		Readytoreceive: true,
// 	}, nil

// }

func Client() {
	log.Println("Client called")
}

package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bbruun/galt/commands"
	minion "github.com/bbruun/galt/proto"
)

type myMinionServer struct {
	minion.UnimplementedMinionServiceServer
}

func (m *myMinionServer) Create(context.Context, *minion.CreateMinion) (*minion.MinionCreateResponse, error) {
	return &minion.MinionCreateResponse{}, nil
}

func (m *myMinionServer) RegisterNewMinion(ctx context.Context, input *minion.CreateMinion) (*minion.MinionCreateResponse, error) {

	a := time.Now().UTC()
	fmt.Printf("On %d RegisterMinion() received: %+v\n", a.UnixNano(), input)
	return &minion.MinionCreateResponse{
		MinionId:       a.UnixMicro(),
		Authenticated:  false,
		Registeredname: input.Name,
	}, nil

}

func (s myMinionServer) GetCommands(srv minion.MinionService_GetCommandsServer) error {
	//ctx := stream.Context()

	done := make(chan bool)
	_ = done

	go func() {
		ctx := srv.Context()
		for {
			// exit if context is done
			// or continue
			select {
			case <-ctx.Done():
				return
			default:
			}
			req, err := srv.Recv()
			if err == io.EOF {
				// return will close stream from server side
				log.Println("client disconnected")
				return
			}
			if err != nil {
				fmt.Printf("receive error: %s\n", err)
				continue
			}
			fmt.Printf("received: %+v\n", req)

		}
	}()
	for {
		command := commands.CmdRun{
			Commandline: "ls -l /",
		}
		msg := minion.CommandToMinion{
			Time:             time.Now().UTC().String(),
			Scheduletask:     false,
			Commandtype:      "cmd.run",
			Marshaledcommand: command.ToString(),
			Timeout:          0,
		}
		if err := srv.SendMsg(&msg); err != nil {
			fmt.Printf("failed to send command to minion, closing server port: %s\n", err)
			srv.Context().Done()
			return fmt.Errorf("client is not connected")
		}
		time.Sleep(2 * time.Second)

	}
	return nil
}

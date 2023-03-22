package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/bbruun/galt/messaging"
	"github.com/bbruun/galt/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
)

type ClientQueue struct {
	Name    string
	Message chan string
}

var Configuration *GaltServerConfig
var ClientQueues *[]ClientQueue

var (
	ConfigFile          string   = "./etc/galt/server.yaml"
	ConfigFileLocations []string = []string{"/etc/galt/server.yaml", "./etc/galt/server.yaml"}
)

func init() {

	log.Println("Loading Galt Server Configuration")
}

func NewGaltServerConfig() *GaltServerConfig {
	c := GaltServerConfig{}
	for x := range ConfigFileLocations {
		if _, err := os.Stat(ConfigFileLocations[x]); err != nil {
			continue
		}
		lf, err := os.ReadFile(ConfigFileLocations[x])
		if err != nil {
			log.Printf("Faild to load configuration file \"%s\": %s\n", ConfigFileLocations[x], err)
			os.Exit(1)
		}
		yaml.Unmarshal(lf, &c)
	}
	return &c
}

func StartServer() {
	// TODO: add foreground/background feature to StartServer
	log.Printf("Starting Salt sever")
	var wg sync.WaitGroup
	wg.Add(1)
	go gRPCServer(wg)
	log.Printf("... server started and is ready to server")
	wg.Wait()

}

type myMinionServer struct {
	proto.UnimplementedMinionServiceServer
}

func registerMinionServices(grpcServer *grpc.Server) {
	reflection.Register(grpcServer) // enable reflection from client side
	service := &myMinionServer{}
	proto.RegisterMinionServiceServer(grpcServer, service)

}

func gRPCServer(wg sync.WaitGroup) error {
	defer wg.Done()

	grpcListener, err := net.Listen("tcp", Configuration.GetListenAddress())
	if err != nil {
		log.Fatalf("error setting up net.Listen(): %s", err)
	}

	grpcServer := grpc.NewServer()
	registerMinionServices(grpcServer)

	err = grpcServer.Serve(grpcListener)
	if err != nil {
		return fmt.Errorf("failed to start gRPC server on port 4505: %s", err)
	}
	return nil
}

// func (m *myMinionServer) Create(context.Context, *proto.CreateMinion) (*proto.MinionCreateResponse, error) {
// 	return &proto.MinionCreateResponse{}, nil
// }

func (m *myMinionServer) RegisterNewMinion(ctx context.Context, input *proto.CreateMinion) (*proto.MinionCreateResponse, error) {
	a := time.Now().UTC()
	fmt.Printf("On %d RegisterMinion() received: %+v\n", a.UnixNano(), input)
	return &proto.MinionCreateResponse{
		MinionId:       a.UnixMicro(),
		Authenticated:  false,
		RegisteredName: input.Name,
	}, nil
}

// func (s *myMinionServer) CmdRun(ctx context.Context, cmdRunSend *proto.CmdRunSend) (*proto.CmdRunResult, error) {
// 	log.Printf("CmdRun: received: %+v\n", cmdRunSend)
// 	//TODO: Send this to all minions that match the list in the CmdRunSend.Minions string
// 	return &proto.CmdRunResult{}, nil
// }

func (s *myMinionServer) CmdRun(fromClient *proto.CmdRunFromClient, toClient proto.MinionService_CmdRunServer) error {
	//ctx := stream.Context()

	done := make(chan bool)
	_ = done

	// Read from Minion
	var mc *messaging.Minions = messaging.MinionStateCollector

	msgch := make(chan any)
	mi := messaging.MinionInfo{
		Name:                  fromClient.Name,
		MessageFromClient:     fromClient.MessageFromClient,
		MessageToClient:       fromClient.MessageToClient,
		CommunicationsChannel: msgch,
		IsConnected:           true,
	}
	mc.AddMinion(&mi)
	log.Printf("CmdRun: received: %+v\n", fromClient)
	// return toclient
	loopIndex := 3
	for {
		b := time.Now().UTC()

		messageToReturn := fmt.Sprintf("The minion %s sent a cmd.run command (id:%d)", mi.Name, b.UnixNano())
		msg := proto.SendCommandToMinion{
			MessageToClient: messageToReturn,
		}

		if err := toClient.SendMsg(&msg); err != nil {
			mi.IsConnected = false
			fmt.Printf("minion %s disconected (%s)\n", mi.Name, err)
			toClient.Context().Done()
			return fmt.Errorf("client %s is not connected", mi.Name)
		}
		time.Sleep(2 * time.Second)
		loopIndex -= 1
		if loopIndex < 0 {
			mi.IsConnected = false
			break
		}
	}
	return nil
}

func (s *myMinionServer) GetCommands(fromClient *proto.CommandFromMinion, toClient proto.MinionService_GetCommandsServer) error {
	//ctx := stream.Context()

	done := make(chan bool)
	_ = done

	// Read from Minion
	var mc *messaging.Minions = messaging.MinionStateCollector

	msgch := make(chan any)
	mi := messaging.MinionInfo{
		Name:                  fromClient.Name,
		MessageFromClient:     fromClient.MessageFromClient,
		MessageToClient:       fromClient.MessageToClient,
		CommunicationsChannel: msgch,
		IsConnected:           true,
	}
	mc.AddMinion(&mi)

	// return toclient
	for {
		b := time.Now().UTC()

		messageToReturn := fmt.Sprintf("The minion %s will see this (id:%d)", mi.Name, b.UnixNano())
		msg := proto.SendCommandToMinion{
			MessageToClient: messageToReturn,
		}

		if err := toClient.SendMsg(&msg); err != nil {
			mi.IsConnected = false
			fmt.Printf("minion %s disconected (%s)\n", mi.Name, err)
			toClient.Context().Done()
			return fmt.Errorf("client %s is not connected", mi.Name)
		}
		time.Sleep(2 * time.Second)

	}

	// Read from Minion
	// go func() {
	// 	ctx := srv.Context()
	// 	for {
	// 		// exit if context is done
	// 		// or continue
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		default:
	// 		}
	// 		req, err := srv.Recv()
	// 		if err == io.EOF {
	// 			// return will close stream from server side
	// 			log.Println("client disconnected")
	// 			return
	// 		}
	// 		if err != nil {
	// 			fmt.Printf("receive error: %s\n", err)
	// 			continue
	// 		}
	// 		fmt.Printf("received: %+v\n", req)

	// 	}
	// }()

	// To client
	// for {
	// 	command := commands.CmdRun{
	// 		Commandline: "ls -l /",
	// 	}
	// 	msg := minion.CommandToMinion{
	// 		Time:             time.Now().UTC().String(),
	// 		Scheduletask:     false,
	// 		Commandtype:      "cmd.run",
	// 		Marshaledcommand: command.ToString(),
	// 		Timeout:          0,
	// 	}
	// 	if err := srv.SendMsg(&msg); err != nil {
	// 		fmt.Printf("failed to send command to minion, closing server port: %s\n", err)
	// 		srv.Context().Done()
	// 		return fmt.Errorf("client is not connected")
	// 	}
	// 	time.Sleep(2 * time.Second)
	// }
	return nil
}

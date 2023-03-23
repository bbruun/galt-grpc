package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
	fmt.Printf("- got: %+v\n", fromClient)
	mi := messaging.MinionInfo{
		Name:                  fromClient.Name,
		MessageFromClient:     fromClient.MessageFromClient,
		MessageFromServer:     fromClient.MessageFromServer,
		CommunicationsChannel: msgch,
		IsConnected:           true,
	}
	mc.AddMinion(&mi)

	targets := strings.Split(fromClient.TargetMinions, ", ")
	command := fromClient.Command

	log.Printf("CmdRun:\n- targets: %s\n- command: %s\n", targets, command)

	//TODO: Actually run the command
	// return output to calling client
	loopIndex := len(targets) - 1
	for {
		// b := time.Now().UTC()

		//TODO: return actual content from above
		messageToReturn := fmt.Sprintf("This should be response from minon %d ... (depending on the order of execution and minion returns of couse)", loopIndex+1)
		msg := proto.SendCommandResultToMinion{
			MessageFromServer: messageToReturn,
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
		MessageFromServer:     fromClient.MessageFromServer,
		CommunicationsChannel: msgch,
		IsConnected:           true,
	}
	mc.AddMinion(&mi)

	//TODO: process received commmand on target minions

	// return results to toclient
	for {
		b := time.Now().UTC()

		msg := proto.SendCommandResultToMinion{
			MessageFromServer: fmt.Sprintf("The minion %s will see this (id:%d)", mi.Name, b.UnixNano()),
		}

		if err := toClient.SendMsg(msg); err != nil {
			mi.IsConnected = false
			fmt.Printf("minion %s disconected (%s)\n", mi.Name, err)
			toClient.Context().Done()
			return fmt.Errorf("client %s is not connected", mi.Name)
		}
		time.Sleep(2 * time.Second)

	}

	return nil
}

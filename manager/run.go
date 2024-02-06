package manager

// protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     bluesdp.proto
//  this is the one that is working
// protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative    bluesdp/bluesdp.proto
import (
	bluesdp_proto "bluesdp/bluerpc"
	"bluesdp/config"
	"bluesdp/database"
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/madflojo/tasks"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	startrpc = &cobra.Command{
		Use:   "sdpserver",
		Short: "Start blue Service Descovery Server",
		Long:  `Start blue Service Regisetry Server`,
		Run: func(cmd *cobra.Command, args []string) {
			RpcServe()
		},
	}

	rpcclient = &cobra.Command{
		Use:   "sdpclient",
		Short: "Start blue Service Regisetry Client",
		Long:  `Start blue Service Regisetry Client`,
		Run: func(cmd *cobra.Command, args []string) {
			RunClient()
		},
	}
)

func RpcServe() {
	lis, err := net.Listen("tcp", "0.0.0.0:"+config.Config("PORT"))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", err, config.Config("PORT"))
	}

	blueserver := bluesdp_proto.BlueSDPServer{}

	grpcServer := grpc.NewServer()

	// Task clearing old data
	scheduler := tasks.New()
	defer scheduler.Stop()

	clear_run, _ := strconv.Atoi(config.Config("CLEAR_RUN_INTERVAL"))
	clear_run = int(clear_run)
	clear_run_interval := time.Millisecond * time.Duration(clear_run)
	if _, err := scheduler.Add(&tasks.Task{
		Interval: time.Duration(clear_run_interval),
		TaskFunc: func() error {
			database.DeleteOutdated()
			return nil
		},
	}); err != nil {
		fmt.Println(err)
	}

	bluesdp_proto.RegisterBlueSDPHeartBeatServiceServer(grpcServer, &blueserver)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
	fmt.Println("Started RPC Server for BLUE")

}

func RpcClient() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.Config("TARGET_HOST")+":"+config.Config("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := bluesdp_proto.NewBlueSDPHeartBeatServiceClient(conn)

	message := bluesdp_proto.BlueHeartBeat{
		ServiceName: "blue-back",
		IpAddress:   config.Config("CONTAINER_IP"),
	}

	_, err3 := c.SayAlive(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err3)
	}

	_, err2 := c.GetList(context.Background(), &bluesdp_proto.GetRegistredServiceList{ServiceName: "blue-back"})
	if err != nil {
		log.Fatalf("Error when calling GetList: %s", err2)
	}

	// log.Printf("Response from Server: %s", response2.Ips)

	// log.Printf("Response from Server: %s", response.Reply)

}

func RunClient() {

	clear_run, _ := strconv.Atoi(config.Config("HEART_BEAT_INTERVAL"))
	clear_run = int(clear_run)
	clear_run_interval := time.Millisecond * time.Duration(clear_run)
	//  Task 2 for testing Make random heartbeat call
	scheduler := tasks.New()
	defer scheduler.Stop()

	if _, err := scheduler.Add(&tasks.Task{
		Interval: clear_run_interval,
		TaskFunc: func() error {
			RpcClient()

			return nil
		},
	}); err != nil {
		fmt.Println(err)

	}
	fmt.Printf("Running Client SDP RPC calls Every  %v", clear_run_interval)
	select {}

}

func MakeRpcCall(ip_address string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(config.Config("TARGET_HOST")+":"+":"+config.Config("PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := bluesdp_proto.NewBlueSDPHeartBeatServiceClient(conn)

	message := bluesdp_proto.BlueHeartBeat{
		ServiceName: "blue-back",
		IpAddress:   ip_address,
	}

	_, err3 := c.SayAlive(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err3)
	}

	_, err2 := c.GetList(context.Background(), &bluesdp_proto.GetRegistredServiceList{ServiceName: "blue-back"})
	if err != nil {
		log.Fatalf("Error when calling GetList: %s", err2)
	}

	// log.Printf("Response from Server: %s", response2.Ips)

	// log.Printf("Response from Server: %s", response.Reply)

}

func init() {
	goBlueCmd.AddCommand(startrpc)
	goBlueCmd.AddCommand(rpcclient)

}

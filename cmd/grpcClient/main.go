package main

import (
  "flag"
  "log"

	"google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "github.com/mtsdalmolin/grpc-protobuf/internal/pb"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
  var opts []grpc.DialOption

  opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
  defer conn.Close()

  client := pb.NewCategoryServiceClient(conn)

  log.Println(client.ListCategories(nil, nil))

}

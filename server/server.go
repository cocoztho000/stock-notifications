/*
stock-notification package sens stock notifications when certain indicators
are triggered Copyright (C) 2019 Tom Cocozzello

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/google/uuid"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"

	storage "github.com/cocoztho000/stock-notifications/etcd"
	pb "github.com/cocoztho000/stock-notifications/stockdata"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

type routeGuideServer struct {
}

// StockRecorder records a stream of stocks
func (s *routeGuideServer) StockRecorder(stream pb.StockData_StockRecorderServer) error {
	var stockCount int32
	var totalStockPrice float32

	// Initialize Etcd
	etcd := storage.NewEtcd()

	// var lastStock *pb.Stock
	startTime := time.Now()
	for {
		stock, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pb.StockSummary{
				StocksReceived: stockCount,
				ElapsedTime:    int32(endTime.Sub(startTime).Seconds()),
			})
		}
		if err != nil {
			return err
		}

		err = etcd.PutWithTimeout(fmt.Sprint("/test/tom/%s/%s", uuid.Must(uuid.NewRandom()), stock.Name), fmt.Sprintf("%f", stock.Price))
		if err != nil {
			fmt.Println("Error writting to etcd: " + err.Error())
		}

		stockCount++
		totalStockPrice += stock.Price
		// if lastStock != nil {
		// 	totalStockPrice += calcStock(lastStock, stock)
		// }
		// lastStock = stock
	}
}

// func calcStock(s1 *pb.Stock, s2 *pb.Stock) float32 {
// 	return s1.Price + s2.Price
// }

func newServer() *routeGuideServer {
	return &routeGuideServer{}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = testdata.Path("server1.pem")
		}
		if *keyFile == "" {
			*keyFile = testdata.Path("server1.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterStockDataServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

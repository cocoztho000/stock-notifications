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
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	sd "github.com/cocoztho000/stock-notifications/stockdata"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

// sendFakeStockData sends fake stock data to the server
func sendFakeStockData(client sd.StockDataClient) {
	var stocks []*sd.Stock = getRandomStockData()

	log.Printf("Sending %d Stock.", len(stocks))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.StockRecorder(ctx)
	if err != nil {
		log.Fatalf("StockRecorder Error: %v", err)
	}
	for _, stock := range stocks {
		if err := stream.Send(stock); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, stock, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("CloseAndRecv Error: %v", err)
	}
	log.Printf("Stock summary: %v", reply)
}

func getRandomStockData() []*sd.Stock {
	var stocks []*sd.Stock

	// Random stock data
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stockCount := int(r.Int31n(100)) + 2
	for i := 0; i < stockCount; i++ {
		stocks = append(stocks, randomStock(r))
	}

	return stocks
}

func randomStock(r *rand.Rand) *sd.Stock {
	long := r.Float32()
	return &sd.Stock{Name: "test", Price: long}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := sd.NewStockDataClient(conn)

	sendFakeStockData(client)
}

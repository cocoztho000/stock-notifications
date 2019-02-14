# stock-notifications
Provide Stock Notifications from Indicators

## Generate grpc Code
```
protoc -I stockdata/ stockdata/stock_data.proto --go_out=plugins=grpc:stockdata
```

## Running Server
```
go run server/server.go -tls=true
```
## Running Client
```
go run aggregator/aggregator.go -tls=true
```

## Getting Started
1. Clone Repo
2. Run `dep ensure` to install dependencies

## Starting a local etcd cluster
```
etcd_version="v3.3.11"
container_name="etcd-gcr-${etcd_version}.7"
rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
  docker rmi quay.io/coreos/etcd:${etcd_version} || true && \
  docker run -d \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name ${container_name} \
  quay.io/coreos/etcd:${etcd_version} \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new

export ETCDCTL_API=3
etcdctl put foo bar
etcdctl get foo
```

## Beginning design:
API
-> Stream of data comes into the api
-> Write the data into etcd
   - `/data/raw/<stock_name>/<date>/price/ = 40.1`
-> Forward Stream of grpc requests to calculation engine to free up api
   - Data: stock names and price points that were written into etcd
Calc engine
-> Stream data comes into the engine 
-> Spawn off goroutine to update all calculations per stock 
   - `/data/algorithm/<stock_name>/<algorithm_name>/<algorithm_timeframe> = 50`
-> Send notification if it has hit

Discovery client
-> Gather stock data
-> stream stock data to the api server


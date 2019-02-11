# stock-notifications
Provide Stock Notifications from Indicators

## Generate grpc Code
```
protoc -I stockdata/ stockdata/stock_data.proto --go_out=plugins=grpc:stockdata
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

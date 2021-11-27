all: client

orders/orders_grpc.pb.go: orders/orders.proto
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative  orders/orders.proto 

client: orders/*.go order_client/*.go
	go build -o client order_client/*.go

clean:
	rm orders/*.go

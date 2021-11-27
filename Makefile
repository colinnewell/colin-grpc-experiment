all: orders/orders_grpc.pb.go

orders/orders_grpc.pb.go: orders/orders.proto
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative  orders/orders.proto 

clean:
	rm orders/*.go

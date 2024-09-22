

# Testing GRPC with evans
1.
> $ evans -r repl
2.
> $ package pb
> $ service OrderService


## To call Create Order:
> $ call CreateOrder

## To call ListOrders:

> $ call ListOrders

## Generating GQL:

```
$ go run github.com/99designs/gqlgen generate

$ go mod tidy
``` 

## Generating GRPC

> $ protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

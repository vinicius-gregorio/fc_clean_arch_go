# Full Cycle - Clean Architecture Challenge. by Vinicius Gregorio


## How to run?
### Run the project with docker-compose
``` $ docker compose up -d```
You should see the 3 containers running with no errors.


![image](https://github.com/user-attachments/assets/5959b23f-dd7a-4f91-8f16-cec82f3004db)

## What port each server responds to?
GraphQL -> 8080

HTTP -> 8000

GRPC -> 50051


## How can I test the servers?

### Testing HTTP server
1. Using VSCode, go to `api/create_order.http`
2. Click on "Send Request" on the request you want to test.
3. You can use any other client, if not VSCode not desired.
4. You should see a response with no errors when creating an order

   
![image](https://github.com/user-attachments/assets/8515762a-24ab-45e5-8337-c2f7d177d77e)


5. You should see a response with no errors when listing orders


![image](https://github.com/user-attachments/assets/22ae58d0-b47d-43d2-bc41-cf4a0cc1f8a8)


### Testing GraphQL server
1. With the APP running, go to [Your 8080 localhost](http://localhost:8080/)
2. Use this mutation to create an order:
```
mutation CreateOrder {
  createOrder(input: { 
    id: "ddd", 
    Price: 100.0, 
    Tax: 10.0 
  }) {
    id
    Price
    Tax
    FinalPrice
  }
}
```
![image](https://github.com/user-attachments/assets/675cd92c-fd48-4aae-9aa6-044002d58b32)

3. Use this mutation to list orders:
```query GetOrders{
  getAllOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
![image](https://github.com/user-attachments/assets/158e10fe-7962-4415-9646-83c80d0ee519)



### Testing GRPC with evans
1.
> $ evans -r repl
2.
> $ package pb
> $ service OrderService


## To call Create Order:
> $ call CreateOrder

![image](https://github.com/user-attachments/assets/61f261d1-eb26-4934-8c06-bab00156da47)

## To call ListOrders:

> $ call ListOrders

![image](https://github.com/user-attachments/assets/a569197c-a264-4d42-bd3a-739b8a815149)



# For other dev or later me: 

### Generating GQL:

```
$ go run github.com/99designs/gqlgen generate

$ go mod tidy
``` 

### Generating GRPC

> $ protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

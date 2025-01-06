package main

import (
	"database/sql"
	"fmt"
	"net"

	"20-CleanArch/configs"
	"20-CleanArch/internal/event/handler"
	"20-CleanArch/internal/infra/grpc/pb"
	"20-CleanArch/internal/infra/grpc/service"
	"20-CleanArch/internal/infra/web/webserver"
	"20-CleanArch/pkg/events"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//region **** RabbitMQ ****
	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("OrderListed", &handler.OrderListedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	//endregion

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db, eventDispatcher)

	//#region **** REST ****
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler("/order", webOrderHandler.Create)
	webserver.AddHandler("/order/list", webOrderHandler.FindAll)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	// cria o webserver em outra thread para não bloquear a execução
	go webserver.Start()
	// #endregion

	//#region **** gRPC ****
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	//listOrderService := service.ListOrderService(*listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer) // ler e processar sua propria informação, usado para o EVANS funcionar

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	// cria o webserver em outra thread para não bloquear a execução
	grpcServer.Serve(lis)
	//#endregion

	//region **** GraphQL ****
	// srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
	// 	CreateOrderUseCase: *createOrderUseCase,
	// }}))
	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	// http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
	//#endregion
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

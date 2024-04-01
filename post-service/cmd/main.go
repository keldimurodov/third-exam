package main

import (
	"net"
	"third-exam/post-service/config"
	pb "third-exam/post-service/genproto/post"
	"third-exam/post-service/pkg/db"
	"third-exam/post-service/pkg/logger"
	"third-exam/post-service/service"
	grpcclient "third-exam/post-service/service/grpc_client"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "third-exam/post-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("grpc client dail error", logger.Error(err))
	}

	// connDB, err, _ := db.ConnectToDB(cfg)
	// if err != nil {
	// 	log.Fatal("sqlx connection to postgres error", logger.Error(err))
	// }

	// postService := service.NewPostService(connDB, log, grpcClient)

	conMongoDB, err := db.ConnectToMongoDB(cfg)
	if err != nil {
		log.Fatal("mongo connection to mongodb error", logger.Error(err))
	}

	postService := service.NewCommentServiceMongo(conMongoDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

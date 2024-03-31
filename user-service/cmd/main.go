package main

import (
	"net"
	"third-exam/user-service/config"
	pbu "third-exam/user-service/genproto/user"
	"third-exam/user-service/pkg/db"
	"third-exam/user-service/pkg/logger"
	"third-exam/user-service/service"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "third-exam/user-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	connRedisDB, err := db.ConnectToRedisDB(cfg)
	if err != nil {
		log.Fatal("connection to Redis error", logger.Error(err))
	}

	userService := service.NewUserService(connDB, connRedisDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbu.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

package main

import (
	"go.uber.org/zap"
	"log"
	"mlm/config"
	"mlm/delivery/http"
	"mlm/logger"
	"mlm/repository/mysql"
	"mlm/repository/mysql/node_repo"
	"mlm/repository/mysql/user_repo"
	"mlm/service/node_svc"
	"mlm/service/user_svc"
)

// @title			Golang
// @version			1.0
// @description		Server API
// @host			localhost:1320
// @BasePath		/
func main() {

	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}

	logger.Logger.Named("main").Info("config", zap.Any("config", config.AppConfig.ZapLogger))

	db, err := mysql.Connect(config.AppConfig.Mysql)
	defer mysql.Close(db)

	if err != nil {
		log.Fatalf("Error mysql connection:%v", err)
	}

	nodeRepository := node_repo.NewNodeRepository(db)
	userRepository := user_repo.NewUserRepository(db)

	userSvc := user_svc.NewUserService(userRepository)
	nodeSrv := node_svc.NewNodeService(nodeRepository, userSvc)

	server := http.New(
		config.AppConfig.Application,
		nodeSrv,
	)

	server.Serve()

}

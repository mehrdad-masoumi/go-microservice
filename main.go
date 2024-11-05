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
	"mlm/validator/node_validator"
	"mlm/validator/user_validator"
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

	// repo
	nodeRepository := node_repo.NewNodeRepository(db)
	userRepository := user_repo.NewUserRepository(db)

	//validator
	nodeValidator := node_validator.NewNodeValidator(nodeRepository)
	userValidator := user_validator.NewUserValidator(userRepository)

	// service
	nodeSrv := node_svc.NewNodeService(nodeRepository, nodeValidator)
	userSvc := user_svc.NewUserService(userRepository, userValidator)

	server := http.New(
		config.AppConfig.Application,
		nodeSrv,
		userSvc,
		nodeValidator,
	)

	server.Serve()

}

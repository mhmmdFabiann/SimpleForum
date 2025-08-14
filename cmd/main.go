package main

import (
	//"database/sql"
	"log"
	"project2/internal/configs"
	"project2/internal/handler/membership"
	"project2/internal/handler/post"
	"project2/pkg/internalsql"

	membershipRepository "project2/internal/repository/memberships"
	membershipSvc "project2/internal/service/memberships"

	"github.com/gin-gonic/gin"

	postRepo "project2/internal/repository/posts"
	postSvc "project2/internal/service/posts"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfFile(
			"config",
		),
		configs.WithConfType(
			"yaml",
		),
	)

	if err != nil {
		log.Fatal("gagal inisiasi config", err)
	}
	cfg = configs.GetConf()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("error initiating database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepository.NewRepository(db)
	membershipService := membershipSvc.NewService(cfg, membershipRepo)

	membershipHandler := membership.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postRepo := postRepo.NewRepository(db)
	postSvc := postSvc.NewService(cfg, postRepo)

	postHandler := post.NewHandler(r, postSvc)
	postHandler.RegisterRoute()

	serverAddress := ":" + cfg.Service.Port
	r.Run(serverAddress)
}